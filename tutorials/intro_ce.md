# Cloud Journey Compute Engine Quickstart

![Cloud Journey](https://walkthroughs.googleusercontent.com/content/gcpquest/title.png "Cloud Journey Tutorial")

<walkthrough-tutorial-url url="https://cloud.google.com/compute/docs/gcpquest/sysintro"></walkthrough-tutorial-url>

## Introduction

<walkthrough-tutorial-duration duration="10"></walkthrough-tutorial-duration>

You've been tasked with creating a very specific VM for Cloud Journey. Don't
worry! Continue with this guide to get it done!

## Project setup

Google Cloud Platform organizes resources into projects. This allows you to
collect all the related resources for a single application in one place. &nbsp;
\
&nbsp; \
*You may have already selected a project in a previews tutorial, if it is in the
box below, then you are all set.* &nbsp; \
&nbsp; \
<walkthrough-project-billing-setup></walkthrough-project-billing-setup>

## Navigate to Compute Engine

Open the [menu][spotlight-console-menu] on the left side of the console.

Then, select the **Compute Engine** section.

<walkthrough-menu-navigation sectionId="COMPUTE_SECTION"></walkthrough-menu-navigation>

*If this is the first time you are using this project, you may have to wait a
few minutes for Compute Engine to be setup for your account.*

## Create a virtual machine instance

Click the [Create instance][spotlight-create-instance] button.

*   Select a [name][spotlight-instance-name] and [zone][spotlight-instance-zone]
    for this instance.

*   In the Machine Type selector, click on
    **[Machine Type][spotlight-machine-type]** dropdown.

*   Select **micro (1 shared vCPU)**

*   Click the [Create][spotlight-submit-create] button to create the instance.

Note: Once the instance is created your billing account will start being charged
according to the GCE pricing. You will remove the instance later to avoid extra
charges.

## VM instances page

While the instance is being created take your time to explore the VM instances
page.

*   At the bottom you can see the [list of your VMs][spotlight-vm-list]
*   At the top you can see a [control panel][spotlight-control-panel] allowing
    you to
    *   Create a new VM instance or an instance group
    *   Start, stop, reset and delete instances

## Connect to your instance

When the VM instance is created, you'll run a web server on the virtual machine.

The [SSH buttons][spotlight-ssh-buttons] in the table will open up an SSH
session to your instance in a separate window.

For this tutorial you will connect using Cloud Shell. Cloud Shell is a built-in
command line tool for the console.

### Wait for the instance creation to finish

The instance creation needs to finish before the tutorial can proceed. The
activity can be tracked by clicking the
[notification menu][spotlight-notification-menu] from the navigation bar at the
top.

### Connect to the instance

Click on the [SSH buttons][spotlight-ssh-buttons] to connect to the machine.
Play around, but you don't have to do anything.

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
