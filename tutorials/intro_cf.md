# Cloud Journey Cloud Functions Quickstart

![Cloud Journey](https://walkthroughs.googleusercontent.com/content/gcpquest/title.png "Cloud Journey Tutorial")

<walkthrough-tutorial-url url="https://cloud.google.com/compute/docs/gcpquest/adintro"></walkthrough-tutorial-url>

## Introduction

<walkthrough-tutorial-duration duration="10"></walkthrough-tutorial-duration>

You've been tasked with creating a Cloud Function that spits out a random
hexadecimal token. Relatively easy to code, and even easier to create the Cloud
Function. This walkthrough will take you through the whole thing.

## Project setup

Google Cloud Platform organizes resources into projects. This allows you to
collect all the related resources for a single application in one place. &nbsp;
\
&nbsp; \
*You may have already selected a project in a previews tutorial, if it is in the
box below, then you are all set.* &nbsp; \
&nbsp; \
<walkthrough-project-billing-setup></walkthrough-project-billing-setup>

## Navigate to Cloud Functions

Open the [menu][spotlight-console-menu] on the left side of the console. &nbsp;
\
&nbsp; \
Then, select the **Cloud Functions** section.

<walkthrough-menu-navigation sectionId="FUNCTIONS_SECTION"></walkthrough-menu-navigation>

## Create a Function

Click on [Create Function][spotlight-function-create]. Change both the
[name][spotlight-function-name] and
[Function to execute][spotlight-function-name] to:

```js
tokenGenerator
```

&nbsp; \
&nbsp; \
Delete all the contents of [index.js][spotlight-code-tab] and then copy and
paste this token generating code into that form field.

```js
exports.tokenGenerator = (req, res) => {
  let message = (Math.random().toString(36)+'00000000000000000').slice(2, 5+2);
  res.status(200).send(message);
};
```

Delete all the contents of [package.json][spotlight-code-tab] and copy and paste
this code:

```js
{
  "name": "token-generator",
  "version": "0.0.1"
}
```

And click on [Create][spotlight-function-create-code] button. &nbsp; \
&nbsp; \
Now we wait... When it is done, click through to the next step of the tutorial.

## Test

Once the function is done spinning up, you can click on
[tokenGenerator][spotlight-function-link]. &nbsp; \
&nbsp; \
Click on the [Testing][spotlight-function-test] tab. &nbsp; \
&nbsp; \
Click on the [Test the function][spotlight-function-test-do] button. &nbsp; \
&nbsp; \
You should get a random token.

## Conclusion

<walkthrough-conclusion-trophy></walkthrough-conclusion-trophy>

You're done!

Go back to the game, and keep questing.

[pricing]: https://cloud.google.com/compute/#compute-engine-pricing
[spotlight-create-instance]: walkthrough://spotlight-pointer?=gce-zero-new-vm,gce-vm-list-new
[spotlight-instance-name]: walkthrough://spotlight-pointer?spotlightId=gce-vm-add-name
[spotlight-instance-zone]: walkthrough://spotlight-pointer?spotlightId=gce-vm-add-zone-select
[spotlight-boot-disk]: walkthrough://spotlight-pointer?cssSelector=vm-set-boot-disk
[spotlight-firewall]: walkthrough://spotlight-pointer?spotlightId=gce-vm-add-firewall
[spotlight-vm-list]: walkthrough://spotlight-pointer?cssSelector=.p6n-checkboxed-table
[spotlight-control-panel]: walkthrough://spotlight-pointer?cssSelector=#p6n-action-bar-container-main
[spotlight-ssh-buttons]: walkthrough://spotlight-pointer?cssSelector=gce-connect-to-instance
[spotlight-notification-menu]: walkthrough://spotlight-pointer?cssSelector=.p6n-notification-dropdown,.cfc-icon-notifications
[spotlight-console-menu]: walkthrough://spotlight-pointer?spotlightId=console-nav-menu
[spotlight-open-devshell]: walkthrough://spotlight-pointer?spotlightId=devshell-activate-button
[spotlight-machine-type]: walkthrough://spotlight-pointer?spotlightId=gce-add-machine-type-select
[spotlight-submit-create]: walkthrough://spotlight-pointer?spotlightId=gce-submit
[spotlight-external-ip]: walkthrough://spotlight-pointer?cssSelector=.p6n-external-link
[spotlight-instance-checkbox]: walkthrough://spotlight-pointer?cssSelector=.p6n-checkbox-form-label
[spotlight-delete-button]: walkthrough://spotlight-pointer?cssSelector=.p6n-icon-delete
[spotlight-machine-type]: walkthrough://spotlight-pointer?spotlightId=gce-add-machine-type
[spotlight-enable-button]: walkthrough://spotlight-pointer?cssSelector=.jfk-button-action
[spotlight-function-execute]: walkthrough://spotlight-pointer?cssSelector=input.jfk-textinput:not(.label-input-label)
[spotlight-function-create]: walkthrough://spotlight-pointer?spotlightId=gcf-list-new
[spotlight-function-create-code]: walkthrough://spotlight-pointer?spotlightId=gce-submit-button
[spotlight-function-name]: walkthrough://spotlight-pointer?cssSelector=.jfk-textinput
[spotlight-code-tab]: walkthrough://spotlight-pointer?cssSelector=.p6n-fullscreen-codemirror
[spotlight-function-link]: walkthrough://spotlight-pointer?cssSelector=.p6n-icon-status
[spotlight-function-test]: walkthrough://spotlight-pointer?cssSelector=.goog-tab:nth-child(4)
[spotlight-function-test-do]: walkthrough://spotlight-pointer?cssSelector=.p6n-loading-button-regular-text
