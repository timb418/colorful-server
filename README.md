# colorful-server
 
This is a simple server that remembers "literature" names for the RGB colors. 
Meaning that for search query rgb(255, 215, 0) it will return "gold".

At a previous iteration I used a SQLite database, but then decided to get rid of it for simplicity.
For HTML page generation it uses html/template package.

Usage: go to http://localhost:8090/colors?rgb=rgb(255,215,0)
Expected result: HTML page with text "gold" and a gold background.
![img.png](img.png)