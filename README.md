# GridPlus Software Engineer Code Test
Welcome to the GridPlus software engineer code test. This test is intended to assess a candidate's skill in solving a realistic software engineering problem which you might come across working at GridPlus.

Below, an open ended software problem based on past work at GridPlus is laid out for you to solve. The requirements are laid out below but how you solve it is completely up to you. Once you are satisfied with your solution, let us know, and we will have you walk us through your solution on a screen share. Completion of the test is intended to take ~3 hours of development time.

You may use any programming language to complete this test. We are most familiar with golang, NodeJS, and Python, but as long as you can walk us through your solution any appropriate language choice is good.

All of the information you will need to design the best solution may not be available or entirely clear in this document, so please feel free to ask us any clarifying questions you may have as you work through the test.

The primary things we are looking for in the assessment are:

- working code that meets the requirements
- clean architecture
- readable code
- clear communication of the design during the walkthrough

## The Problem
For this code test, you will build a mock version of the software update system which runs on the Lattice1's GCE.

### Background
The Lattice1 is a consumer hardware device that will typically be plugged in at a customer's home and, after initial setup, regularly connected to the internet.

Internally, it is divided between two distinct hardware environments, the GCE, or General Compute Environment, and the SCE, or Secure Compute Environment (also known as the HSM.)

The HSM is physically segregated from the outside world, communicating only with the GCE over a secure serial interface we call the "mailbox."

The GCE is responsible for all communication with the outside world over the internet, therefore the responsibility of updating the HSM's software falls on the GCE.

### Requirements
When run, your program should contact another GridPlus service, the release-catalog-api, to detect available updates, and then, when appropriate, should "update" the relevant applications. This updater program should be able to manage updates for both the GCE and HSM.

In this code test, since we will not have actual software to update, applying an "update" should be mocked by simply printing a message to the terminal including the contents of the software update binary at the time when the update would normally take place.

If there are updates available for both the HSM and the GCE at the same time, the HSM update should always be applied first, in order to manage backward-compatibility for the mailbox protocol.

The updater program should have some way of persistently storing the current versions of the GCE and HSM. If the current version of an application is greater than or equal to the available version, an update should not be triggered.

Software versions are handled using a simplified semantic versioning scheme, where versions are given as `Major.Minor.Fix` there are no extensions or other special rules involved in our versioning scheme.

Update targets listed in the release catalog may have a `PrereqUpdates` field. This field identifies update versions which are prerequisites before installing the version listed. If the current version is older than any of the prerequisite updates, these must be applied sequentially before the top level update can be applied safely. This is in order to safely manage software updates that may involve migration of persistent data which would break if previous migrations have not yet been applied.


### Querying Available Updates

Available updates should be detected by querying another GridPlus service, the release-catalog-api. When new versions of software are ready to be released, their information is published to this service to inform Lattices in the field that there is a software update available.

An example of response data from this endpoint looks like the following:
```
{
  "targets": [
    {
      "AppCode": "HSM",
      "TargetVersion": "0.9.8",
      "downloadURL": "https://release-catalog-api.gridpl.us/download/repository/lattice-updates/hsm/hsm-0.9.8-production.bin.sig",
      "DownloadSize": 326740,
      "ArtifactSig": ""
    },
    {},
    {
      "AppCode": "GCE",
      "TargetVersion": "0.48.6",
      "downloadURL": "https://release-catalog-api.gridpl.us/download/repository/lattice-updates/gce/lattice-gce-0.48.6-production.bin.sig",
      "DownloadSize": 11010260,
      "ArtifactSig": ""
    }
  ]
}
```

A mock version of this service has been set up specifically for this code test at `interview-release-catalog-api.staging-gridpl.us/update`

In order to request update targets from this service, you must send a POST request to this HTTPS endpoint with an UpdateRequest JSON payload, structured as follows:
```
[{"AppCode": "GCE", "CurrentVersion": "Major.Minor.Fix"}, {"AppCode": "HSM", "CurrentVersion": "Major.Minor.Fix"}]
```

This endpoint is protected by a basic authentication challenge with the following credentials.
```
username: lattice1
password: codetest
```

## Deliverables

The following should be delivered as a part of this code test.

1. An executable program which when run, detects updates and applies them according to the requirements given above.
2. Some way to change the stored current versions in between runs of the program.
3. Basic instructions for running the program on our machines. Ideally, we should be able to run your program to test it out before you walk us through your design.

Please push your commits directly to this repository as you complete them. This gives us an easy way to check in on your progress and to review your work when completed. Just let us know through slack when you feel your test is complete and we'll review and schedule a walkthrough with you.

Happy coding!






