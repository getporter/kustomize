# Welcome to the `porter-kustomize` mixin

## What is `porter-kustomize`? 

### Let's start with what is [Porter.sh](https://porter.sh/)....

!!! quote
    ![Porter Logo](https://porter.sh/images/porter-logo.png){: style="height:125px; width:150px display: block; margin: 0 auto"}
    *A Friendly Cloud Installer for Cloud Native Application Bundles*
    
    When we deploy to the cloud, most of us aren’t dealing with just a single cloud provider or toolchain. The simplest
    of applications today need a load balancer, SSL certificate, persistent file storage, DNS, and somewhere in there
    is your application. One app is installed with Helm, another with the cloud provider’s cli and it is all glued
    together with magic bash scripts.
    
    That is a lot to figure out! 😅  
    
    Porter is a cloud installer based on the Cloud Native Application Bundle (CNAB) spec that helps you manage
    everything together in a single bundle, focusing on what you know best: your application.
    

Porter ships with a number of builtin plugins known as `mixins` that facilitate
deployments of CNAB bundles.

!!! info
    You should read and undertsand the documentation on the [Porter.sh](https://porter.sh/) website and
    [GitHub repo](https://github.com/deislabs/porter) along with the [CNAB Website](https://cnab.io/)
    and [CNAB reference documentation](https://github.com/deislabs/cnab-spec) to fully understand the necessary
    concepts.

### What is [Kustomize](https://kustomize.io/)?

!!! quote
    ![Kustomize](./images/3rdParty/kustomize.png){: style="height:450px; width:400px display: block; margin: 0 auto"}
    
    *Kubernetes Native Configuration Management*
    
    Kustomize introduces a template-free way to customize application configuration that simplifies the use of
    off-the-shelf applications.

### So to what is `porter-kustomize`?

The `porter-kustomize` mixin helps extend [`porter.sh`]((https://porter.sh/)) in order to support 
[kustomize.io](https://kustomize.io/) which provides Kubernetes Native Configuration Management from porter.
