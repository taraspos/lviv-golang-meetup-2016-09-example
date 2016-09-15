# AWS

## Create Golang AppService as AWS Elastic Beanstalk Application

### Amazon Console [Elastic Beanstalk page](http://console.aws.amazon.com/elasticbeanstalk/)

- Create new app

    ![](static/aws/1-create-app.png)

  - Set service name and description
  - Create web server environment
  - Choose Go language(current version 1.5)
  - Select application:
    - Sample App
    - From S3 bucket
    - Upload zip file (more details on the next slides)
  - Configure deployment preferences
  - Configure environment(name, url)
  - Additional resources(Network, DB connection)
  - Instance configuration(size, SSH key, notifications, health-check…)
  - Configure tags
  - VPC
  - Permissios(IAM role)


### Elastic Beanstalk CLI

- Install eb-cli with:

        pip install awsebcli

- Init eb project with:

        eb init

- Follow next prompts to:
  - set region
  - set aws credentials
  - choose application name
  - choose programming language environment
  - choose programming language version
  - setup ssh

- Create environment with:

        eb create {env name}

- Deploy new application version:

        eb deploy {env name}




## Deploy your app

- Open deployment options:

    *App Services* **->** *{AppServiceName}* **->** *Deployment Options*

- Configure deployment source
  - Choose one of supported deployment sources
  - Authorize deployment source;
  - Choose Repo and Branch to deploy

    ![](static/azure/3-deploy-app.png)

- (Optional) Create deployment slots for different environments(dev/stage/prod).

    ![](static/azure/4-deployment-slots.png)


## Additional tweaking

- Application environment variables can be set via:

   *App Services* **->** *{WebServiceName}* **->** *Application settings*
- In order to connect WebService to the cloud VMs, we need to:
  - Create Virtual Network
  - Attach VMs network interfaces to that VirtualNetwork
  - Create Virtual Network Gateway for WebService
  - Attach Virtual Network Gateway to the WebService
- Connection to databases or storage blobs could be done via: 

   *App Services* **->** *{WebServiceName}* **->** *Data Connections*
- Exploring WebService guts could be done via:

  *App Services* **->** *{WebServiceName}* **->** *Advanced Tools*

  or

  https://{WebServiceName}.scm.azurewebsites.net/ 
