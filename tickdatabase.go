package tickdatabase

//Key is our generic name for accessing data
type Key string


//TickDBEntry  - Wrapper to do what I want to do
type TickDBEntry map[Key]string
//TTickDB contains the latest and greatest time value
type TTickDB struct {
  Current TickDBEntry
  //This design might bite me in the butt once additional collector threads are thrown in.
  Historical map[Key][]string
}
