# Ray's Recipe Catalog
## A catalog of my favorite recipes

Ray's Receipe Catalog is a Go application with a simple frontend to manage a catalog of recipes I like to cook.

I build this webapp because I wanted to learn Go, get experience with DynamoDB and Terraform, and keep track of recipes I like to cook without ads.

## Built with
- Go - for the backend
- DynamoDB - for the database (AWS NoSQL)
- Terraform - for deploying the DynamoDB instance
- HTML - for the frontend

## Features
- Add and update recipes with a name, cuisine, ingredient list, cooking instructions, source, and cook time
- Search for what you're in the mood for by name, cuisine, and/or ingredient list

## Add a Recipe
Add recipes with all available information. In this case, I'm adding a simple scrambled eggs recipe from foodnetwork.com.
![](https://github.com/raytighe/recipe_catalog/blob/main/img/add-recipe.gif)

## Search Recipes
Search by keyword within each respective field. Search results are the intersection of multiple keyword results. For example, searching for "Eggs" by recipe name will return all recipes with eggs in the title. Searching by "Asian" cuisine and "chicken" in the ingredients will return all Asian recipes that require chicken. 
![](https://github.com/raytighe/recipe_catalog/blob/main/img/search.gif)

## Update Recipes

## Installation