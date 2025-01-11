## Attribute Server

**Purpose:**

Supply additional information needed in an Attribute based access control system.

**How it works:**

Register the app in azure ad. Every other backend application is a client to this app.

For each client backend app:
1. The client backend will send request for user attributes to the microservice along with a validation token/or this client app is added to a security group.
2. The microservice will query the attribute tables after validating the token and will respond back to the client backend
3. The client microservice will use the attributes to build ABAC logic

This microservice provides the data to be checked by ABAC rules implemented in OPA service.

## Design

### Authorization

#### Azure
This app will be registered with Azure AD. We will need to expose the api, create app roles. The consuming apps will have to be added to the appropriate roles and scopes. Admin consent has to be provided to the consuming apps   