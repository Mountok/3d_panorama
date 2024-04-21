var currentImage;


    var path = window.location;
    currentImage = path.href.split("=")[1]
    // console.log(currentImage)

    pannellum.viewer('panorama', {
        "type": "equirectangular",
        "panorama": `http://localhost:8080/images?id=${currentImage}`,
        "autoLoad": true,
    }); 




    document.addEventListener('DOMContentLoaded', function() {
        var button = document.querySelector('button');
        button.addEventListener('click', function() {
          window.location.href = `https://go-uni-3d-pano.netlify.app/index.html`;
        });
      });