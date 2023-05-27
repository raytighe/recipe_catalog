# Ray's Recipe Catalog
## A catalog of my favorite recipes

Ray's Receipe Catalog is a Go application with a simple front end to manage a catalog of recipes I look to cook.

I build this webapp because I wanted to learn Go, get experience with DynamoDB and Terraform, and keep track of recipes I like to cook without ads.

## Built with
- Go - for the backend
- DynamoDB - for the database (AWS NoSQL)
- Terraform - for deploying the DynamoDB instance
- HTML - for the frontend

## Features
- Add and update recipes with a name, cuisine, ingredient list, cooking instructions, source, and cook time
- Search for what you're in the mood for by name, cuisine, and ingredient list

## Add a Recipe
![](https://github.com/raytighe/recipe_catalog/blob/main/img/add-recipe.gif)

## Search Recipes

## Update Recipes

## Installation

Dillinger requires [Node.js](https://nodejs.org/) v10+ to run.

Install the dependencies and devDependencies and start the server.

```sh
cd dillinger
npm i
node app
```

For production environments...

```sh
npm install --production
NODE_ENV=production node app
```