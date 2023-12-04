package main
import ("fmt")

func main() {

    // Name your variable for the works//
var password int 
var choice1 string
var choice2 string

fmt.Println("welcome to kuda bank ")
fmt.Println("please Enter Your Password ")
fmt.scan( &password)
// Request for password
if (password ==9237)
  // clear digite with getchar
fmt.Println("Your balance is 0.00")
fmt.Println(" 1. Transfar money")
fmt.Println(" 2. Add money")
fmt.Println(" 3. Buy Airtime")
fmt.Println(" 4. Internet")
fmt.Println(" 5. Electicity") 
//Ask the user to make is chioce
fmt.Println(" select number")
fmt.scan(&choice1)
}
switch (choice1){
case (1): fmt.Println("Chose your bank")
          fmt.Println("1. first Bank")
          fmt.Println("2. Eco Bank")
          fmt.println("3. Access bank")
          fmt.println("4. Other bank")
          fmt.scan("%d", &choice2)
 }

if  (choice2 == 1){
  fmt.Println(" Account number")
  }
else if (choice2 == 2){
  fmt.Println(" Account number for Eco")

}
else if (choice2 == 3)
 {
 fmt.Println(" Account number for Access")

 }
else if (choice2 == 4)
{
  fmt.Println("choose your bank")
          }
          
          break;
 case (2): 
           { 
              fmt.Println("1. Overdraft")
              fmt.Println("2. Send Money")
              fmt.Println("3. USSD")
              fmt.Println("4. Card")
              fmt.Println("5. Casch Deposit")
              fmt.scan ( &choice2)
              if  (choice2 == 1)
{
  fmt.Println("How much")
}
else if (choice2 == 2)
{
  fmt.Println(" From where to where")

}
else if (choice2 == 3)
{
 fmt.Println(" Put Your code")

}
else if (choice2 == 4)
{
  fmt.Println(" verve")
          }
    break
    case(3): fmt.Println(" Chose an amount")
             fmt.Println("1. #100")
             fmt.Println("2. #200")
             fmt.Println("3. #500")
             fmt.Println("4. #1000")
             fmt.Println("5. #2000")
             fmt.Println("6. #3000")
             fmt.scan(&choice1)
             if  (choice3 == 1)
 {
  fmt.Println("1. Buy")
 }
else if (choice3 == 2)
 {
  fmt.Println("2. Buy")

 }
else if (choice3 == 3)
 {
 fmt.Println("2.Buy")

 }
else if (choice3 =)
         fmt.Println("2.Buy ")
    break
    case(4): fmt.Println("1. select service provider")
             fmt.Println("2. SPECTRANNET")
             fmt.Println("3. SWIFT NG")
             fmt.Println("4. SMILE NG")
             fmt.Println("5. IPNX NG")
             fmt.Println("6. MTN DATA")
     break
     case(5): fmt.Println(" select service provider")
               fmt.Println("1. AEDC NG")
               fmt.Println("2. APLE NG (ABA POWER)")
               fmt.rintln("3. BEDC NG")
default:
fmt.Println("Invalid entry, please try again later.")
    break;
}
else 
{
fmt.Println("Invalid entry, please try again later.")
}

return 0

 }






//  #include <stdio.h>
// #include <stdlib.h>
// int main()
// {
//     // Name your variable for the works//
// int code;
// int choice1;
// int choice2;

// printf("\nwelcome to kuda bank \n");
// printf("\nplease Enter Your Password \n");
// scanf("%d", &code);
// // Request for password
// if (code ==9237)
// {
//   // clear digite with getchar
// getchar();
// printf("\n Your balance is 0.00");
// printf("\n 1. Transfar money");
// printf("\n 2. Add money");
// printf("\n 3. Buy Airtime");
// printf("\n 4. Internet");
// printf("\n 5. Electicity");
// //Ask the user to make is chioce
// printf("\n select number");
// scanf ("%d",&choice1);
// switch (choice1)
// {
// case (1): printf("\nChose your bank");
//           printf("\n1. first Bank");
//           printf("\n2. Eco Bank");
//           printf("\n3. Access bank");
//           printf("\n4. Other bank");
//           scanf("%d", &choice2);

// if  (choice2 == 1)
// {
//   printf("\n Account number");
// }
// else if (choice2 == 2)
// {
//   printf("\n Account number for Eco");

// }
// else if (choice2 == 3)
// {
//  printf("\n Account number for Access");

// }
// else if (choice2 == 4)
// {
//   printf("\n choose your bank");
//           }
          
//           break;
//  case (2): printf("\n1. Overdraft");
//               printf("\n2. Send Money");
//               printf("\n3. USSD");
//               printf("\n4. Card");
//               printf("\n5. Casch Deposit");
//               scanf("%d", &choice2);
//               if  (choice2 == 1)
// {
//   printf("\n How much");
// }
// else if (choice2 == 2)
// {
//   printf("\n From where to where");

// }
// else if (choice2 == 3)
// {
//  printf("\n Put Your code");

// }
// else if (choice2 == 4)
// {
//   printf("\n verve");
//           }
//     break;
//     case(3): printf("\n Chose an amount");
//              printf("\n1. #100");
//              printf("\n2. #200");
//              printf("\n3. #500");
//              printf("\n4. #1000");
//              printf("\n5. #2000");
//              printf("\n6. #3000");
//              scanf("%.2lf",&choice1)
//              if  (choice3 == 1)
// {
//   printf("\n1. Buy");
// }
// else if (choice3 == 2)
// {
//   printf("\n2. Buy");

// }
// else if (choice3 == 3)
// {
//  printf("\n2.Buy");

// }
// else if (choice3 =)
// printf("2.Buy ")
//     break;
//     case(4): printf("\n1. select service provider");
//              printf("\n2. SPECTRANNET");
//              printf("\n3. SWIFT NG");
//              printf("\n4. SMILE NG");
//              printf("\n5. IPNX NG");
//              printf("\n6. MTN DATA");
//      break;
//      case(5): printf("\n select service provider");
//                printf("\n1. AEDC NG");
//                printf("\n2. APLE NG (ABA POWER)");
//                printf("\n3. BEDC NG");
// default:
// printf("Invalid entry, please try again later.\n");
//     break;
// }
// }
// else 
// {
// printf("Invalid entry, please try again later.\n");
// }

// return 0;
// }