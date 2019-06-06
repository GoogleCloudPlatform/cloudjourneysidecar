## Cloud Journey BigQuery Quickstart

![Cloud Journey](https://walkthroughs.googleusercontent.com/content/gcpquest/title.png "Cloud Journey Tutorial")

<walkthrough-tutorial-url url="https://cloud.google.com/compute/docs/gcpquest/bdintro"></walkthrough-tutorial-url>

## Introduction

<walkthrough-tutorial-duration duration="10"></walkthrough-tutorial-duration>

You've been tasked with querying a very specific question using BigQuery: what
was the second most common name for individuals born in the US in 1976. Don't
worry, this tutorial will walk you through it.

## Project setup

Google Cloud Platform organizes resources into projects. This allows you to
collect all the related resources for a single application in one place. &nbsp;
\
&nbsp; \
*You may have already selected a project in a previews tutorial, if it is in the
box below, then you are all set.* &nbsp; \
&nbsp; \
<walkthrough-project-billing-setup permissions="compute.instances.create"></walkthrough-project-billing-setup>

## Navigate to BigQuery

Open the [menu][spotlight-console-menu] on the left side of the console.

Then, select the **BigQuery** section.

<walkthrough-menu-navigation sectionId="BIGQUERY_SECTION"></walkthrough-menu-navigation>

## Ask the right question

The question you were asked was about demographic data, specifically popular
names in a given year. We can use one of public datasets on Big Query to answer
that. &nbsp; 

Copy and paste this query into the [Query Editor][spotlight-query].

```sql
SELECT name, SUM(number) as total
FROM `bigquery-public-data.usa_names.usa_1910_2013`
WHERE year = 1976
GROUP BY name, gender
ORDER BY total
DESC
LIMIT 10
```

Then click the [Run Query][spotlight-run] button. &nbsp; \
&nbsp; \
When it is done, you should see that *Jennifer* is the second most common name
for individuals born in 1976.

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
[spotlight-query]: walkthrough://spotlight-pointer?cssSelector=.p6n-code-mirror-editor
[spotlight-run]: walkthrough://spotlight-pointer?cssSelector=.p6n-split-button
