const header = document.querySelector("header");

window.onscroll = function() {
  if (window.scrollY > 75) {
    header.classList.add("scroll");
  } else {
    header.classList.remove("scroll");
  }
};


// code for search bar filter using JSON data.
fetch('search.json')
  .then(response => response.json())
  .then(data => {
    const dataList = document.getElementById('searchList');
    const inputSearch = document.getElementById('search');
    const highlight = document.getElementById('highlight');

    data.forEach(item => {
      const option = document.createElement('option');
      option.value = item.doctor + " - " + item.specialty;
      dataList.appendChild(option);
    });

    inputSearch.addEventListener('input', function() {
      const inputValue = inputSearch.value.toLowerCase();
      let highlightText = '';

      data.forEach(item => {
        const optionText = (item.doctor + " - " + item.specialty).toLowerCase();
        if (optionText.startsWith(inputValue)) {
          highlightText = item.doctor.substring(0, inputValue.length);
        }
      });

      highlight.innerText = highlightText;
    });

    inputSearch.addEventListener('keydown', function(event) {
      if (event.key === "Tab" && dataList.firstChild) {
        inputSearch.value = dataList.firstChild.value;
      }
    });
  })
  .catch(error => console.error('Error fetching data:', error));


// showing all doctors and minimizing them
document.getElementById('showAll').addEventListener('click', function() {
  const specialists = document.querySelectorAll('.specialist');
  const button = document.getElementById('showAll');
  const areAllShown = specialists[12].classList.contains('show'); // Check if specialists beyond 12 are shown

  if (!areAllShown) {
    // Show all specialists and change button text to 'Show Less'
    specialists.forEach((specialist, index) => {
      if (index >= 12) specialist.classList.add('show');
    });
    button.innerHTML = 'Show Less <i class="ph-fill ph-circles-four text-3xl"></i>';
  } else {
    // Hide specialists beyond the first 12 and change button text to 'Show More'
    specialists.forEach((specialist, index) => {
      if (index >= 12) specialist.classList.remove('show');
    });
    button.innerHTML = 'Show More <i class="ph-fill ph-circles-three-plus text-3xl"></i>';
  }
});

document.getElementById('viewAll').addEventListener('click', function() {
  const diseases = document.querySelectorAll('.diseases');
  const button = document.getElementById('viewAll');
  const areAllShown = diseases[17].classList.contains('show'); // Check if diseases beyond 17 are shown

  if (!areAllShown) {
    // Show all diseases and change button text to 'Show Less'
    diseases.forEach((disease, index) => {
      if (index >= 17) disease.classList.add('show');
    });
    button.innerHTML = 'Show Less <i class="ph-fill ph-circles-four text-3xl"></i>';
  } else {
    // Hide diseases beyond the first 12 and change button text to 'Show More'
    diseases.forEach((disease, index) => {
      if (index >= 17) disease.classList.remove('show');
    });
    button.innerHTML = 'Show More <i class="ph-fill ph-circles-three-plus text-3xl"></i>';
  }
});


// WhatsApp form submission
document
  .querySelector("form")
  .addEventListener("submit", function(event) {
    event.preventDefault();

    let name = document.getElementById("name").value;
    let email = document.getElementById("email").value;
    let phone = document.getElementById("phone").value;
    let app = document.getElementById("app").value;
    let diagnosis = document.getElementById("diagnosis").value;
    let contact = "+256758185721";

    let customer = name.replace(/(^\w{1})|(\s+\w{1})/g, (letter) =>
      letter.toUpperCase()
    );

    let encodedMessage = encodeURIComponent(
      "Name: " +
      customer +
      "\n" +
      "Email: " +
      email +
      "\n" +
      "Phone: " +
      phone +
      "\n" +
      "Preferred App: " +
      app +
      "\n" +
      "Diagnosis: " +
      diagnosis
    );

    let link;

    // Check if user is on a mobile device
    if (
      /Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(
        navigator.userAgent
      )
    ) {
      link = `whatsapp://send?phone=${contact}&text=${encodedMessage}`;
    } else {
      // Desktop device
      link = `https://web.whatsapp.com/send?phone=${contact}&text=${encodedMessage}`;
    }

    window.open(link, "_blank");
  });