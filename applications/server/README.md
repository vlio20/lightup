# DB:
docker run --name lightup -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=root -e MONGO_INITDB_ROOT_PASSWORD=root  -d mongo


# dreate new module
1. cd modules
2. cp <EXISTING_MODULE> <NEW_MODULE>
3. for f in $(find <NEW_MODULE> -name '*.go'); do mv "$f" "$(echo "$f" | sed s/<EXISTING_MODULE>/<NEW_MODULE>/)"; done
4. (UpperCamelCase) find ./<NEW_MODULE> -type f -exec sed -i '' -e 's/<EXISTING_MODULE>/<NEW_MODULE>/g' {} +
5. Same as 4 but with snake_case
6. Same as 4 but with lowerCase