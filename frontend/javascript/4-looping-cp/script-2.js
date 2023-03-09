// Mengembalikan pola * yang menurun di setiap barisnya
//
// contoh:
// baris = 5
// hasil:
// *****
// ****
// ***
// **
// *

// Masukan jumlah baris
const n = 10;

// TODO: answer here

let m = n;

for (let i = 0; i < n; i++) {
  for (let i = 0; i < m; i++) {
    process.stdout.write("*");
  }
  m -= 1;
  console.log("");
}
