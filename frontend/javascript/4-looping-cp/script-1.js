// Mengembalikan urutan bilangan, dan mengetahui bilangan kelipatan 3 dan kelipatan 5
// contoh:
// baris = 5
// hasil:
// 1
// 2
// 3 merupakan kelipatan 3
// 4
// 5 merupakan kelipatan 5
//
// catatam: ingat, 15 merupakan kelipatan 3 dan 5.

// Masukan jumlah bilangan yang ingin dicek, iterasi dari angka 1
//const counter = parseInt(prompt("Masukan jumlah bilangan yang ingin dicek: "));

// TODO: answer here

const counter = 30;
const isMultiple3 = "merupakan kelipatan 3";
const isMultiple5 = "merupakan kelipatan 5";

for (var num = 1; num <= counter; num++) {
  if (num % 3 == 0 && num % 5 == 0) {
    console.log(num + " " + "merupakan kelipatan 3 dan 5");
  } else if (num % 3 === 0) {
    console.log(num + isMultiple3);
  } else if (num % 5 === 0) {
    console.log(num + isMultiple5);
  } else {
    console.log(num);
  }
}
