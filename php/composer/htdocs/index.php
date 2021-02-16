<!DOCTYPE html>
<html>
  <head>
    <title>Powered By Paketo Buildpacks</title>
  </head>
  <body>
    <img style="display: block; margin-left: auto; margin-right: auto; width: 50%;" src="https://paketo.io/images/paketo-logo-full-color.png"></img>
<?php
  // https://getcomposer.org/doc/01-basic-usage.md#autoloading
  // This is how you autoload composer packages
  require '../vendor/autoload.php';

  $dotenv = Dotenv\Dotenv::createImmutable(__DIR__);
  $dotenv->load();
  $projectName = $_ENV['PROJECT_NAME'];
  echo "<p style='text-align: center'>Powered By " . $projectName . " Buildpacks</p>"
?>
  </body>
</html>
