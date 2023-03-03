#### 
```
To implement this Oauth2.0 for any application like Google, Facebook, Github we first need to register our application on their portal, from where we will get client id anf cliend secret key.
```

#### Google
```
Go to the Google Cloud Console.
Create a new project or select an existing project.
In the left sidebar, select APIs & Services > Credentials.
Click the Create Credentials button and select OAuth client ID.
Choose your application type (Web application, Mobile application, or Desktop application) and enter the required information.
After creating your OAuth 2.0 client ID, you will be provided with a client ID and secret.
```

#### Facebook
```
Go to the Facebook Developers website.
Create a new app or select an existing app.
In the left sidebar, select Products > Facebook Login > Settings.
Under Client OAuth Settings, enter your redirect URL and allowed domains.
After saving your settings, you will be provided with a client ID and secret.
```

#### GitHub
```
Go to your GitHub Developer settings.
Click the New OAuth App button.
Enter your application name, homepage URL, and callback URL.
After registering your application, you will be provided with a client ID and secret.
```


#### 
```
After register your application on the portal then you need to feed your secret keys in this application in config file.

Note:- For github the code implementation is little diffrent so in this code we will see an index.html file where we need to configure our redirect url in which you have to pass your client id .
```
