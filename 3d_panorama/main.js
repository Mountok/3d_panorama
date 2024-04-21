document.addEventListener('DOMContentLoaded', function () {
  var select = document.querySelector('.classic.images');
  var button = document.querySelector('button');
  var selectedOption
  button.addEventListener('click', function () {
    selectedOption = select.options[select.selectedIndex].value;
    //   console.log('Selected option:', selectedOption);
  });

  var button = document.querySelector('button');
  
  button.addEventListener('click', function () {
    if (selectedOption == "none") {
      return 
    }
    window.location.href = `https://go-uni-3d-pano.netlify.app/view.html?image=${selectedOption}`;
  });
});

let imagesData;
let ownerData = {};

window.onload = async () => {
  await fetch("http://localhost:8080/images-all")
    .then(response => response.json())
    .then(data => {
      imagesData = data.data;
      console.log(imagesData)
    })
    .catch(error => {
      console.log("ошибка")
    })

  let selectElement = document.querySelector(".classic.images");

  for (let i of imagesData) {

    // let optionElement = document.createElement("option");
    // optionElement.value = i.image_url;
    // optionElement.text = i.image_name;
    // selectElement.add(optionElement)
    // console.log(i)
    ownerData[i.image_owner] = true
  }

  selectElement = document.querySelector(".classic.owners");
  for (let i in ownerData) {
    let optionElement = document.createElement("option");
    optionElement.value = i;
    optionElement.text = i;
    selectElement.add(optionElement)
  }
}

selectElement = document.querySelector(".classic.owners");

selectElement.addEventListener("change", function () {
  let currOwner = this.value

  let imagesSelectElement = document.querySelector(".classic.images");
  let imagesOptionElements = imagesSelectElement.querySelectorAll("option")

  for (let i = 1; i < imagesOptionElements.length; i++) {
    imagesSelectElement.removeChild(imagesOptionElements[i])
  }


  for (let i of imagesData) {
    if (i.image_owner == currOwner) {
      let optionElement = document.createElement("option");
      optionElement.value = i.image_url;
      optionElement.text = i.image_name;
      imagesSelectElement.add(optionElement)
    }
  }

})
