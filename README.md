# Instagram-Backend-API

The task is to develop a basic version of aInstagram. You are only required to develop the API for the system. Below are the details.

This API system has following functions:

1) Create a User

URL is ‘/users'

          -User ID
          -Name
          -Email ID
          -Password --> Password encrypted
          
2) Get a user using id

URL is  ‘/users/<id>’

3) Create a Post
         
URL is ‘/posts'

        -User ID 
        -Post ID
        -Caption
        -Image URL
        -Posted Time Stamp

4) Get a post using id
          
URL is ‘/posts/<id>'
          
5) List of all posts of a user --> Not running
          
URL is ‘/posts/users/<id>'
