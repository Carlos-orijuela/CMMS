

$(document).ready(function(){
    $('#example').DataTable();
});

// var tableData = [
//     { name: "John Doe", code: "john@example.com", asset: "1234567890", vendor: "1234567890", location: "1234567890" , barcode: "1234567890"},
//     // Add more data as needed
//   ];
  
//   var currentPage = 1;
//   var rowsPerPage = 10; // Adjust the number of rows per page as needed
//   var totalPages = Math.ceil(tableData.length / rowsPerPage);
  
//   function displayTableData() {
//     var startIndex = (currentPage - 1) * rowsPerPage;
//     var endIndex = startIndex + rowsPerPage;
//     var tableBody = document.querySelector("#my-table tbody");
//     tableBody.innerHTML = "";
  
//     for (var i = startIndex; i < endIndex; i++) {
//       if (i >= tableData.length) {
//         break;
//       }
  
//       var rowData = tableData[i];
//       var row = document.createElement("tr");
//       row.innerHTML = `<td>${rowData.name}</td>
//                        <td>${rowData.code}</td>
//                        <td>${rowData.asset}</td>
                       
//                        <td>${rowData.location}</td>
//                        <td>${rowData.barcode}</td>
//                        <td><button class="custom-button" id="openModalBtnassetupdate"><span class="material-icons-outlined" title="Update">create</span> </button>
//                        <button class="custom-button" id="openModalBtndelete"><span class="material-icons-outlined" title="Delete">delete</span> </button>
//                        <button class="custom-button" id="openModalBtn"><span class="material-icons-outlined" title="Settings">settings</span> </button></td>`;
//       tableBody.appendChild(row);
//     }
//   }
  
//   function setupPagination() {
//     var paginationContainer = document.querySelector("#pagination");
//     paginationContainer.innerHTML = "";
  
//     for (var i = 1; i <= totalPages; i++) {
//       var listItem = document.createElement("li");
//       listItem.textContent = i;
  
//       if (i === currentPage) {
//         listItem.classList.add("active");
//       }
  
//       listItem.addEventListener("click", function () {
//         currentPage = parseInt(this.textContent);
//         displayTableData();
//         setupPagination();
//       });
  
//       paginationContainer.appendChild(listItem);
//     }
//   }
  
//   displayTableData();
//   setupPagination();
