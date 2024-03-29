<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8" />
    <meta
      name="viewport"
      content="width=device-width, initial-scale=1, shrink-to-fit=no"
    />
    <meta name="description" content="" />
    <meta name="author" content="" />

    <title>{{.Title}}</title>

    <!-- Bootstrap core CSS -->
    <link href="{{.RelativePath}}vendor/bootstrap/css/bootstrap.min.css" rel="stylesheet" />

    <!-- Custom fonts for this template -->
    <link
      href="{{.RelativePath}}vendor/fontawesome-free/css/all.min.css"
      rel="stylesheet"
      type="text/css"
    />
    <link
      href="https://fonts.googleapis.com/css?family=Lora:400,700,400italic,700italic"
      rel="stylesheet"
      type="text/css"
    />
    <link
      href="https://fonts.googleapis.com/css?family=Open+Sans:300italic,400italic,600italic,700italic,800italic,400,300,600,700,800"
      rel="stylesheet"
      type="text/css"
    />

    <!-- Global site tag (gtag.js) - Google Analytics -->
    <script async src="https://www.googletagmanager.com/gtag/js?id=G-P7FV7SH5VW"></script>
    <script>
      window.dataLayer = window.dataLayer || [];
      function gtag(){dataLayer.push(arguments);}
      gtag('js', new Date());

      gtag('config', 'G-P7FV7SH5VW');
    </script>

    <!-- Custom styles for this template -->
    <link href="{{.RelativePath}}css/clean-blog.min.css" rel="stylesheet" />

    <!-- favicon -->
    <link rel="apple-touch-icon" sizes="180x180" href="{{.RelativePath}}apple-touch-icon.png">
    <link rel="icon" type="image/png" sizes="32x32" href="{{.RelativePath}}favicon-32x32.png">
    <link rel="icon" type="image/png" sizes="16x16" href="{{.RelativePath}}favicon-16x16.png">
    <link rel="manifest" href="{{.RelativePath}}site.webmanifest">
  </head>

  <body>
    <!-- Navigation -->
    <nav class="navbar navbar-expand-lg navbar-light fixed-top" id="mainNav">
      <div class="container">
        <a class="navbar-brand" href="{{.RelativePath}}index.html">drincruz.com</a>
        <button
          class="navbar-toggler navbar-toggler-right"
          type="button"
          data-toggle="collapse"
          data-target="#navbarResponsive"
          aria-controls="navbarResponsive"
          aria-expanded="false"
          aria-label="Toggle navigation"
        >
          Menu
          <i class="fas fa-bars"></i>
        </button>
        <div class="collapse navbar-collapse" id="navbarResponsive">
          <ul class="navbar-nav ml-auto">
            <li class="nav-item">
              <a class="nav-link" href="{{.RelativePath}}index.html">Home</a>
            </li>
            <li class="nav-item">
              <a class="nav-link" href="{{.RelativePath}}about.html">About</a>
            </li>
            <li class="nav-item">
              <a class="nav-link" href="{{.RelativePath}}archive.html">Archive</a>
            </li>
          </ul>
        </div>
      </div>
    </nav>

    <!-- Page Header -->
    <header class="masthead" style="background-image: url('{{.RelativePath}}img/about-bg.jpg')">
      <div class="overlay"></div>
      <div class="container">
        <div class="row">
          <div class="col-lg-8 col-md-10 mx-auto">
            <div class="page-heading">
              <h1>{{.ContentTitle}}</h1>
              <span class="subheading">{{.ContentSubtitle}}</span>
            </div>
          </div>
        </div>
      </div>
    </header>
