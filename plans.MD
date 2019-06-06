# Cloud Journey Plans


**All tasks depend on the first one**


|Track|Quest|Name|Description|Dependencies|Status|
|---|---|---|---|---|---|
|**General** |1|Install App Engine Application |Install an AppEngine application|None|**Done** | 
|**System** |1|Create a Micro VM |Spin up a f1-micro instance.|None|**Done** | 
|**System** |2|Create a Disk |Create a simple disk to attach to a VM|System 1|**Done** | 
|**System** |3|Create a Managed Instance Group| Create a managed instance group|None|Wanted | 
|**System** |4|Create a Load Balancer|Create a Load Balancer to expose the managed instance group going.|System 3|Wanted | 
|**System** |5|Launch a VM that keeps AutoML model awake |Build a VM that keeps an AutoMl model up and running. |Data 5|Wanted | 
|**App Dev** |1|Create a Cloud Function |Write a simple cloud function|None|**Done** | 
|**App Dev** |2|Create a Cloud Run Application |Publish a simple Cloud Run application|None|Wanted | 
|**App Dev** |3|Attached Cloud Function to Storage Bucket |Do something everytime a files is uploaded to a storage bucket. |Data 2|Wanted  | 
|**App Dev** |4|Launch a static site |Publish a simple HTML app to Cloud Storage.|None|Wanted  | 
|**App Dev** |5|Create an Applicaiton that exposes an Auto ML model |Build a front end that makes an AutoML model public|None|Wanted  | 
|**Data** |1|Do a BigQuery query |Query one of the public datasets on BiqQuery|None|**Done** | 
|**Data** |2|Create a Cloud Storage Bucket |Create a Cloud Storage Bucket|None|Wanted | 
|**Data** |3|Create a Cloud SQL Instance |Create a Cloud SQL Instance |Data 2|Wanted  | 
|**Data** |4|oad a large dataset into BigQuery |Upload and query a large dataset in BigQuery|None|Wanted  | 
|**Data** |5|Create an Auto ML model |Use Auto ML Model to analyze images|None|Wanted  | 





