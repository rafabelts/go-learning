package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
    assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin){
        t.Helper()
        got := wallet.Balance()

        if got != want {
            t.Errorf("got %s want %s", got, want)
        }

    }

    assertErrorWithdrawing := func(t testing.TB, got error, want string){
        t.Helper()
        if got == nil {
            t.Fatal("didn't get an error but wanted one")
        }

        if got.Error() != want {
            t.Errorf("got %q, want %q", got, want)
        } 

    }

    t.Run("deposit", func(t *testing.T){
        wallet := Wallet{}
        wallet.Deposit(Bitcoin(10))
        want := Bitcoin(10)
        assertBalance(t, wallet, want)
    })

    t.Run("withdraw", func(t *testing.T){
        wallet := Wallet{balance: Bitcoin(20)}
        result := wallet.Withdraw(Bitcoin(5))
        want := Bitcoin(15)

        fmt.Println(result)
        assertBalance(t, wallet, want)
    })

    t.Run("withdraw insufficient funds", func(t *testing.T){
        startingBalance := Bitcoin(20)
        wallet := Wallet{startingBalance}
        err := wallet.Withdraw(Bitcoin(100))
        assertErrorWithdrawing(t, err, ErrInsufficientFunds)
        assertBalance(t, wallet, startingBalance)
    })
}
