# Pollywog

[![pipeline status](https://gitlab.com/christianpflugradt/pollywog/badges/main/pipeline.svg)](https://gitlab.com/christianpflugradt/pollywog/-/commits/main) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

Pollywog is an app for carrying out polls. Participants in a poll can add new options to a poll 
and vote for options in the poll. Every poll is private and participants are invited by e-mail
via a personalized link.

Pollywog is a server application providing a REST interface to create polls and participating in them.
The recommended front end to use Pollywog with is [Motivote](https://gitlab.com/christianpflugradt/motivote).

## Deploy an instance of Pollywog ##

Each version of Pollywog, represented by a git tag, is built in Gitlab CI and can be downloaded from there
as long as the build artefacts are available. However the pipeline uses the latest version of GLIBC available
in the environment and thus it is likely that the compiled artefact will not run in your environment.
For that reason the recommended way to obtain a binary of Pollywog is to compile it yourself.

Pollywog expects exactly one argument to be passed on start up. That argument is the file name of the 
configuration file which contains basic information required for Pollywog to run.
An [exemplary configuration file](https://gitlab.com/christianpflugradt/pollywog/-/blob/master/example-config.yml)
is checked into the codebase as part of the project.

Assuming the binary is named `pollywog` and the configuration file `pollywog.yml` then you start Pollywog as follows:
`./pollywog pollywog.yml`

The recommended way to run Pollywog in production is packed as a Docker container.

## Compile Pollywog from source ##

As you can see in the [gitlab-ci.yml](https://gitlab.com/christianpflugradt/pollywog/-/blob/master/.gitlab-ci.yml)
of the project, there are basically four to five steps required to compile Pollywog yourself:
 1. Install [Go](https://golang.org/)
 2. Add the project to your GOPATH
 3. Download the dependencies
 4. Set the version
 5. Run the build command

### Add the project to your GOPATH ###

How you do it is up to you. If you're not familiar with Go and only use it for Pollywog
I'd recommend to set up a symbolic link:
 * `cd ~/go/src`
 * `ln -s /path-to-pollywog-project/src pollywog`

I assume the same can be achieved under Windows using a shortcut.

### Download the dependencies ###

Pollywog depends on the following libraries (install them via go get)
* github.com/go-sql-driver/mysql
* github.com/mattn/go-sqlite3
* gopkg.in/yaml.v2

### Set the version & Run the build command ###

As you can see in Gitlab CI the Pollywog version is set as a Go constant using git describe:

`echo "package model; const Version = \"`git describe --tags`\"" > src/domain/model/version.go`

Setting the version is optional, but it will be revealed via Pollywog's API
and therefore it is a good idea to write the correct version into the binary. The version should
be equal to the git tag you checked out. Eh, you're on master? Nah, better check out the
[latest tag](https://gitlab.com/christianpflugradt/pollywog/-/tags) instead, you know `git checkout tag-name`.

Now once that is done simply run the build command from the project root
and it will produce a `pollywog` binary file: `go build src/pollywog.go`

## Configuration ##

Pollywog uses YAML as configuration format. The following properties are currently available
as part of the configuration:

**client.baseurl**: This URL plus the personalized secret are sent to all participants in a poll.
If you're using Motivote this property should be the location of your Motivote instance
plus the fragment URL char (#) as Motivote reads the secret from the fragment. Review the exemplary
configuration file how a link to Motivote can look like.

**server.port**: This is the port Pollywog will listen to for incoming requests.

**server.admintoken**: Deprecated and will be removed in Pollywog 2.x.x

**server.admintokens**: Holds a list of administrative tokens 
consisting of the following three properties described below: **user**, **token** and **whitelist**

**server.admintokens[].user**: Optional user name associated with an admin token.
If present it will be mentioned in the invitational mail to a poll.

**server.admintokens[].token**: This token secures administrative actions like creating a poll.
Pollywog uses [SHA-512/256](https://en.wikipedia.org/wiki/SHA-2) to hash tokens. 
This configuration field contains the hashed token, so that the unhashed token, that must be present
in the Authorization header when creating a poll, cannot be derived from the configuration easily.
If multiple persons are privileged to create polls it is recommended that each of them have their own admintoken
and do not share their unhashed tokens with each other.

[This tool](https://emn178.github.io/online-tools/sha512_256.html) can hash your token, alternatively 
[just use Go ](https://gitlab.com/christianpflugradt/pollywog/-/blob/master/src/domain/service/authorization.go)
(the function Hash()).

**server.admintokens[].whitelist**: If empty no whitelist will be applied. Otherwise: 
Holds a list of strings all participant mail addresses must match
for an admin to invite them to a poll. Be aware that neither regular expressions nor classic wildcards are supported,
instead a whitelist entry matches a mail address if the mail address ends with that whitelist entry.
If the whitelist is used to restrict who an admin can invite to a poll. It is recommended to either
add complete domains or specific mail addresses to the whitelist. For example having "@mail-example.org"
and "john.doe@example-mail.com" in the whitelist will allow an admin to invite everyone with a mail address
in the domain "@mail-example.org" as well as John Doe from the example-mail.com domain.

**database.driver**: Pollywog uses the `database/sql` interface of Go. The recommended database and driver
for production is `mysql`. For development, I have used `sqlite` but due to its shortcomings it is not
recommended to be used in a productive system with concurrent access.

**database.dataSourceName**: For mysql the proper string is `user:password@pollywog` assuming the database
runs on localhost with default port, the user name being `user`, the password being `password` and the name
of the Pollywog database / schema is `pollywog`.

**smtp**: Smtp configuration is required to send the invitations to all participants when creating a poll.
An own e-mail server is not necessarily required to use Pollywog. A free e-mail address is enough 
and in this case you would simply enter the same values as when using an e-mail client such as Thunderbird.
If your smtp server does not require authentication, leave the password empty. The smtp user will be used as
sender in any case.

## Application programming interface ##

### Administrative actions ###

There is currently one administrative action:
 * create a poll
 
#### A word about security ####

Administrative actions involve all things that could be misused and create damage. For instance
anonymous creation of new polls could result in thousands of polls being created and thus
thousands of e-mails being sent in a short time, possibly including recipients you never wanted to send any e-mails to.
In the worst case your e-mail server may become [blacklisted](https://en.wikipedia.org/wiki/Domain_Name_System-based_Blackhole_List).

In order to prevent that make sure neither the *server.admintoken* nor its public counterpart ever become public.
It is also recommended to change your admintoken regularly. As the configuration file is not part of the binary
and can be updated any time this is very easy to achieve and you could even write a script that changes the token
on a regular basis and notifies all admins about their new public key.
 
#### Create a poll ####

At this point Motivote does not offer a user interface for administrative actions,
so the recommended way is to use a tool for http requests like [Postman](https://www.postman.com/)
or [Insomnia](https://insomnia.rest/).

**endpoint**: The endpoint for creating new polls is `/poll`

**http method**: The http method for creating new polls is `POST`

**http header**: There are two header fields you should set in your request:
 * Header: `Content-Type`, Value: `application/json`
 * Header: `Authorization`, Value: `your unhashed admintoken`
 
**request body**: The following is a legitimate body to create a new poll. Make sure the request body is valid
[JSON](https://www.json.org/json-en.html).
```
{
	"title": "my first poll",
	"description": "this is my first poll, please add any options and vote for them",
	"deadline": "2025-12-31",
	"participants": [
		{
			"name": "Christian",
			"mail": "christian@thisdomaindoesnotexist.org"
		},
		{
			"name": "Linus",
			"mail": "linus@thisdomaindoesnotexist.org"
		}
	]
}
```

### Client development ###

Motivote is the official client for Pollywog but due to the open API it is very easy to write your own client.
As poll participants are invited via e-mail it is currently recommended that your client can be called via a URL.
This does not necessarily mean that your client must be web based. Instead you may register a custom 
[URI schema](https://en.wikipedia.org/wiki/Uniform_Resource_Identifier) for your users linked to your application.

The entry point for all http requests is 
[this file](https://gitlab.com/christianpflugradt/pollywog/-/blob/master/src/web/server.go).
Here you can see which http methods are accepted, what request body is expected and which domain functions are called.

Let's have a look at an example, shall we?
```
func Serve() {
	http.HandleFunc("/poll", multiPoll)
	http.HandleFunc("/options", postOptions)
	http.HandleFunc("/votes", postVotes)
```
This is an excerpt of the main function `Serve()` which exposes all end points that Pollywog will listen to.
As you can see the `/options` endpoint is mapped to the `postOptions` function.

```
func postOptions(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var request representation.OptionsRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err == nil {
			pollId, participantId := service.ResolveParticipant(r.Header.Get("Authorization"))
			if pollId != -1 {
				options := transformer.TransformOptionsRequest(pollId, participantId, request)
				valid := service.UpdatePollOptions(pollId, participantId, options)
				if valid {
					getPoll(w, r)
				} else {
					w.WriteHeader(http.StatusUnprocessableEntity)
				}
			} else {
				w.WriteHeader(http.StatusUnauthorized)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		util.HandleError(util.ErrorLogEvent{ Function: "web.postOptions", Error: err })
	} else if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
```
This function equals to the `/options` endpoint. It only accepts POST requests but it tolerates OPTIONS requests
and simply returns OK for them. All other requests will be responded to with the `Method Not Allowed` status code.

Given we receive a POST request the next step is to decode the request body into an OptionsRequest.
As you can see in the `presentation` package the request body is expected to look like this:
``` 
type OptionsRequest struct {
	KeepOptions []int `json:"keep_options"`
	CreateOptions []string `json:"create_options"`
}
```
The body is expected to contain a list of integers (denoting the participant's options for the poll 
that should be kept) and a list of strings (denoting the new options the participant wants to add).

If the json deserialization fails the request will be responded to with the `Bad Request` status code.

Next in the function the participant will be resolved and for that the `Authorization` header will be analysed.
The value expected in the header is actually the personalized part of the Url the participant received via e-mail
when the poll was created.

If that step fails the request will be responded to with the `Unauthorized` status code.

Now if everything went fine until now the domain function `UpdatePollOptions` will be called.

The domain function returns a boolean indicating the success of the operation. Any constraint violations
such as invalid ids in the keep_options property will result in `valid` being false
and the request will be responded to with the `Unprocessable Entity` status code.

If the request is deemed valid, Pollywog will return `getPoll()` which results in a `OK` response
and a response body containing the current state of the poll after the request has been processed.

## FAQ ##

### Where are the unit tests? ###

The 1.0.0 version of Pollywog along with the 1.0.0 version of Motivote has been written
within three weeks in my spare time, meaning after my regular work and on weekends.

As my goal was to get a production-ready version running within such a short amount of time
some things had to be neglected. 

I have not yet decided if I will add unit tests subsequently. I will however certainly do that once
the codebase continues to grow and manual testing becomes a bigger effort than writing automatic tests.

### Why do I have to compile Pollywog myself? ###

I'm working on a solution to provide binaries permanently and with a better support 
for different operating systems and architectures but it is not ready yet.

### I want feature XYZ ###

Open an issue on Gitlab and I will consider implementing that feature.
