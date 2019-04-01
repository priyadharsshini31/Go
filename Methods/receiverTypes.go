package main

import(
   "fmt"
)


type flowers struct{
	name,season string
}

func(f flowers) changeName(newName string){
	f.name = newName
  //Inside the method you can see the name change
	fmt.Println("Am inside the function and the new name is", f.name) // O/P :lilly
}

func(f *flowers) changeSeason(newSeason string){
        f.season = newSeason
        //Inside the method you can see the season change
        fmt.Println("Am inside the function and the new season is", f.season) // O/P :winter
}

func main(){
	flr := flowers{"jasmine","summer"}
	flr.changeName("lilly");
  //Value as receiver and the change is not visible to the caller, (ie) outside the method
	fmt.Println("The name during value change is", flr.name) // O/P: jasmine
	fmt.Println("The season before change is", flr.season) // O/P: summer
  //Pointer as receiver and the change is visible to the caller, (ie) outside the method
  flr.changeSeason("winter");
	fmt.Println("The Season during pointer chnage is", flr.season) // O/P: winter
}
	
	
