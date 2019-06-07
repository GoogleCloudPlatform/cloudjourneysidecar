# Cloud Journey Project Quickstart

![Cloud Journey](https://walkthroughs.googleusercontent.com/content/gcpquest/title.png "Cloud Journey Tutorial")

<walkthrough-tutorial-url
url="https://cloud.google.com/compute/docs/gcpquest/intro_project">
 </walkthrough-tutorial-url>

## Introduction

<walkthrough-tutorial-duration duration="10"></walkthrough-tutorial-duration>

Before you can can play the game, you have to make sure that you have a Google
Cloud account ready and a project to run the game. **If you are already a Google
Cloud user, you must create a new project to play the game.** \
&nbsp; \
This tutorial will instruct you to install an App Engine Application into your
project. **If you already have an App Engine application it will replace it.**
That's why you should use a clean project. \
&nbsp; \
Additionally, the game will ask you to perform technical tasks in the Cloud
Console. All of these have been designed to use free options. &nbsp; \
To create a new project for the game follow below. Please note: Limit project
names to **30 characters** or less. This is a limitation of the engine behind
Cloud Journey, and not of GCP. \
&nbsp; \
<walkthrough-project-billing-setup></walkthrough-project-billing-setup>

## Record Project ID

**Please *Copy the ID* of the project you created. You will need this in the
game** \
&nbsp; \
You can see it by clicking the [Project Selector][spotlight-purview-switcher]

## Install Game Helper Application

The game uses an App Engine application to make various progress checks. You
must install this to progress in the game. **If you already have an App Engine
application it will replace it.**

Open Cloud Shell by clicking
<walkthrough-cloud-shell-icon></walkthrough-cloud-shell-icon> in the navigation
bar at the top of the console.

&nbsp;

A Cloud Shell session opens inside a new frame at the bottom of the console and
displays a command-line prompt. It can take a few seconds for the shell session
to be initialized. \
&nbsp; \
When Cloud Shell is ready enter the following:

```bash
git clone https://github.com/GoogleCloudPlatform/cloudjourneysidecar.git
```

When done enter the following:

```bash
cd cloudjourneysidecar && make
```

Once the App Engine application is installed you should see a success message.
&nbsp;

## Conclusion

<walkthrough-conclusion-trophy></walkthrough-conclusion-trophy>

You're done!

Go back to the game, and keep questing.

[spotlight-purview-switcher]: walkthrough://spotlight-pointer?spotlightId=purview-switcher
