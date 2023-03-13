package entities

type Client struct {
	shopEntityMetadata

	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	Identification   string `json:"identification"`
	SocialWork       string `json:"social_work"`
	SocialWorkNumber string `json:"social_work_number"`
	FinalConsumer    bool   `json:"final_consumer"`
	Phone            string `json:"phone"`
	Email            string `json:"email"`
	Addess           string `json:"addess"`
	Locality         string `json:"locality"`
	Neighborhood     string `json:"neighborhood"`
	ZipCode          string `json:"zipcode"`
	AditionalInfo    string `json:"aditional_info"`
}
