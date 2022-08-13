Screechr

Step 1:
Create a local directory and enter into it using "cd xx" command
git clone git@github.com:pierbin/screechr.git

Step 2:
Start running the program using the following command
    go run cmd/main.go cmd/wire_gen.go
This program will listen port:80
Two user profiles are created in memory database (sqlite3), the IDs of which are 1 and 2

Step 3:
Open Postman web and create your own workspace
https://web.postman.co/workspace

Step 4:
Find "screechr.postman_collection.json" in the root directory and import into postman (collection v2.1)

Step 5:
You can see 6 APIs as follows:
1. get-a-profile (GET): http://127.0.0.1/profile/1
There is a key called "Authorization" in postman Headers, the value is "xYz123" for first profile, and "aBc123" for second profile
The query param is 1, which is the profile Id. You also can put another query param "2".
You can send the request with Authorization header and query param to get a profile

2. update-profile (POST): http://127.0.0.1/profile/1
Need the same Authorization header "xYz123" to update entire profile fields
The body json data has been created. You can find them in the "Body" tab
You can send the request with Authorization header and body raw Json data

3. create-screech (POST):http://127.0.0.1/screech?creatorid=1
It needs the query param creatorid, both id "1" and "2" can be used (two profiles are available)
The body json data has been created. You can find them in the "Body" tab
You can send the request with body raw Json data and creatorid to create a new screech
You can repeat this step with different body data to create more screeches

4. get-a-screech (GET): http://127.0.0.1/screech/1
You can find a new created screech with id=1 after create-screech
Now you can get the result by sending this request

5. update-screech (POST): http://127.0.0.1/screech/1
The body json data has been created. You can find them in the "Body" tab
You can send the request with body raw Json data and screech id to update a screech's text

6. getscreechlist (GET): http://127.0.0.1/screechlist?creatorid=2&order=asc&size=1
This is to get a list of screeches.
Therea are 3 query params: creatorid, order and size


Step 6: Test
in the root directory, type "cd internal/controller" and run "go test -v"