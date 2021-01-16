package server

const addItemHtml string = `
<!DOCTYPE html>

<html>
<head>
<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, minimum-scale=1, user-scalable=no, minimal-ui">
<style>

/*

menu

*/
/* Add a black background color to the top navigation */
.topnav {
  background-color: rgb(51, 51, 51);
  overflow: hidden;
  margin-left: -8px;
  margin-right: -8px;
  margin-top:-8px;
  
}

/* Style the links inside the navigation bar */
.topnav a {
  float: left;
  color: #f2f2f2;
  text-align: center;
  padding: 14px 16px;
  text-decoration: none;
  font-size: 17px;
}

/* Change the color of links on hover */
.topnav a:hover {
  background-color: rgb(232, 225, 247);
  color: black;
}

/* Add a color to the active/current link */
.topnav a.active {
  background-color: rgb(97, 76, 174);
  color: white;
}

/* end menu*/




html, body {
    
    font-family: Arial, Helvetica, sans-serif;
    height: 100%;
    text-align: center;
}


#container {
    
    margin-bottom: 2em;
    min-height: 100%;
    overflow: auto;
    padding: 0 1em;
    text-align: justify;
}
#footer {
    bottom: 0;
    color: #707070;
    height: 2em;
    left: 0;
    position: relative;
    font-size: small;
    width:100%;
}

.info {
    font-size: small;
    color: #707070;
    
}

.error {
  color:red;
  font-size: small;
  text-align: center;
}

input[type=text] {
  width:100%;
  height: 25px;
  margin-top: 5px;
  margin-bottom: 5px;
}
label {
  font-weight: bold;
}

input[type=submit] {
  background-color: rgb(97, 76, 174); /* Green */
  border: none;
  color: white;
  padding: 10px 26px;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 12px;
  margin: 2px 2px;
  cursor: pointer;
  -webkit-appearance: none;
   border-radius: 0;
}


</style>
<title >Admin</title>
</head>
<div class="topnav">
    <a  href="/">Home</a>
    <a class="active" >Add Item</a>
    <a href="/about">About</a>
  </div> 
  <body >
    <div id="container">
    <h1>Add Item</h1>
      <p class="error" >{{.Data}}</p>
    <form  action="/add" method="POST">
        <label for="iname">Item Name</label><br>
        <input type="text" id="iname" name="iname" value=""><br>
        <label for="url">Url</label><br>
        <input type="text" id="url" name="url" value=""> <br><br>
        <input type="submit" value="Submit">
      </form>

      <p class="info">Items added are not persisted. Add items to settings.json to persist.</p>
    </div>
      <div id="footer">Caleb McCarthy 2021</div>
</body>

</html>


<script>


</script>

`