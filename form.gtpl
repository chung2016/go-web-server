<html>
<head>
<title></title>
</head>
<body>
<form action="/form" method="post">
    username:<input type="text" name="username"><br>
    password:<input type="password" name="password"><br>
    <input type="hidden" name="token" value="{{.}}">
    <input type="submit" value="login">
</form>
</body>
</html>