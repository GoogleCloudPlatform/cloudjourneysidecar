# Cloud Journey Compute Engine Disk Quickstart

![Cloud Journey](https://walkthroughs.googleusercontent.com/content/gcpquest/title.png "Cloud Journey Tutorial")

<walkthrough-tutorial-url url="https://cloud.google.com/compute/docs/gcpquest/adintro"></walkthrough-tutorial-url>

## Introduction

<walkthrough-tutorial-duration duration="10"></walkthrough-tutorial-duration>

You've been tasked with adding an extra disk to a Compute Engine instance that
you set up earlier.  This tutorial will walk you through that.  

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
Select the **Compute Engine** section, then choose **Disks**

<walkthrough-menu-navigation sectionId="COMPUTE_SECTION"></walkthrough-menu-navigation>



## Create a Disk

Click on [Create Disk][spotlight-create-disk]. 

Give it a [name][spotlight-function-name].

Make sure that the `region` and `zone` match the VM that you want to add the disk to. 

You can leave the rest at their default settings. 

And click on [Create][spotlight-submit-disk] button. &nbsp; \
&nbsp; \
Now we wait... When it is done, click through to the next step of the tutorial.

## Attach  Disk

Click on [VM Instances][spotlight-instances]. 

Click on the instance you created in previous quest. 

Click on `Edit`

Scroll down to `Attach Existing Disk` button. 
 
Select disk that you created in previous step. 

Then click `Done`.

Then scroll down and click on `Save` button.


## Conclusion

<walkthrough-conclusion-trophy></walkthrough-conclusion-trophy>

You're done!

Go back to the game, and keep questing.

[spotlight-create-disk]: walkthrough://spotlight-pointer?cssSelector=.cfc-icon-disk-new
[spotlight-submit-disk]: walkthrough://spotlight-pointer?spotlightId=gce-submit-button
[spotlight-instances]: walkthrough://spotlight-pointer?cssSelector=cfc-icon-instance