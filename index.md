# LeadFuze Developer Quiz

## 1. Which one would you prefer and why?

Both of the functions return the same result for a two dimensional integer array (1024x1024). Which one would you prefer? List pros and/or cons of both functions.

```
#define SIZE 1024

int sumA(int a[SIZE][SIZE])
{
  int sum = 0;
  for(int y = 0; y < SIZE; y++)
    for(int x = 0; x < SIZE; x++)
      sum += a[x][y];

  return sum;
}

int sumB(int *a)
{
  int sum[4] = {0, 0, 0, 0};

  for( int i = 0; i < SIZE*SIZE; i += 4 )
  {
      sum[0] += a[i+0];
      sum[1] += a[i+1];
      sum[2] += a[i+2];
      sum[3] += a[i+3];
  }

  return sum[0] + sum[1] + sum[2] + sum[3];
}
```

## 2. What controversies do you see in the following API example?

User object API doc

```
GET /users
POST /users/new
POST /users/:id/update
POST /users/:id/rename
POST /users/:id/update-timezone
DELETE /users/delete?id=:id
```

Here are some examples of the behavior:

```
POST /users/new HTTP/1.1
{
  "name": "Cthulhu"
}

HTTP/1.1 200
"Error: Username already exists"
```

## 3. What problems related to password security can you see in the following example?

A web service stores user information in a database and uses passwords for authentication. Here's how the user password storing and authentication is implemented in ruby (the actual data storage and retrieval is outside the scope of the example):

```
require 'digest'

class User

  # Use salted passwords
  PASSWORD_SALT="trustno1"

  # Stored password hash will be accessible through user.hashed_password
  attr_accessor :hashed_password

  # Authenticates user against given password and returns true
  # the password matches the stored one
  def verify_password(password)
    if hashed_password.nil? || password.nil?
      false
    else
      User.hash_password(password) == hashed_password
    end
  end

  # Changes user's password
  def change_password(new_password)
    self.hashed_password = User.hash_password(new_password.to_s)
  end

  # Hashes the input with salt
  def self.hash_password(password)
    Digest::MD5.hexdigest(password + PASSWORD_SALT)
  end
end
```

## 4. How does the following function violate good design principles?

```
function getDepositHistorySum(user) {
  var deposits = user.transactions.history.deposits;
  var sum = 0;
  for (var i = 0; i < deposits.length; i += 1) {
    sum += deposits[i].amount;
  }
  return sum;
}
```

## 5. What problems does this method have?

```
Account.prototype.increaseBalance = function(amount, isCredit) {
  if (!isCredit) {
    this.debitBalance += amount;
  } else {
    this.creditBalance += amount;
  }
};
```

## 6. List four typical solutions to optimize database query performance

Database queries are getting slow when the database size increases. What are some of the typical solutions to improve performance?

## 7. Coding task 1

Create a brand new Go repo in your GitHub account and implement the following code:

```
package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Println(v)
	}
    /*
        this should print something like this:
            1
            2
            3
            4
            5
            6
            7
            8
        make sure that you close both channels and program should exit successfully at the end.
    */
}

func merge(a, b <-chan int) <-chan int {
    // this should take a and b and return a new channel which will send all values from both a and b
}

func asChan(vs ...int) <-chan int {
    // this should reutrn a channel and send `vs` values randomly to it.
}
```

Return URL of your repo.

## 8. Coding task 2

Go to the following JSFiddle: https://jsfiddle.net/leadfuze/61sr3jdk/. See the code comments for the assignment. Remember to click save and return the url for your fiddle.

## 9. Coding task 3

Clone https://github.com/leadfuze/developer-test-1. Check the code comments for assignment. Push the completed code to your cloned repo and return the cloned repository URL.