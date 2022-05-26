package types

type Roaster struct {
    Id int
    Name string
    CountryId int
    Address string
}

type User struct {
    Id int
    UserName string
    Email string
    FirstName string
    LastName string
    Password []byte
    Salt []byte
}
type Bean struct {
    Id int
    RoasterId int
    Name string
    Description string
    VarietyId int
    ProcessId int
    CountryId int
    RegionId int
    Origin string
}

type Review struct {
    Id int
    ReviewerId int
    Rating int
    Text string
    BeanId int
    RoastDate string
}

