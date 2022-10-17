package inflector_test

import (
	"fmt"

	"github.com/PrimerAI/inflector"
)

func ExamplePluralize() {
	fmt.Println(inflector.Pluralize("post"))
	fmt.Println(inflector.Pluralize("octopus"))
	fmt.Println(inflector.Pluralize("sheep"))
	fmt.Println(inflector.Pluralize("words"))
	fmt.Println(inflector.Pluralize("CamelOctopus"))
	// Output:
	// posts
	// octopi
	// sheep
	// words
	// CamelOctopi
}

func ExampleSingularize() {
	fmt.Println(inflector.Singularize("posts"))
	fmt.Println(inflector.Singularize("octopi"))
	fmt.Println(inflector.Singularize("sheep"))
	fmt.Println(inflector.Singularize("word"))
	fmt.Println(inflector.Singularize("CamelOctopi"))
	// Output:
	// post
	// octopus
	// sheep
	// word
	// CamelOctopus
}

func ExampleCamelize() {
	fmt.Println(inflector.Camelize("my_account"))
	fmt.Println(inflector.Camelize("user-profile"))
	fmt.Println(inflector.Camelize("ssl_error"))
	fmt.Println(inflector.Camelize("http_connection_timeout"))
	fmt.Println(inflector.Camelize("restful_controller"))
	fmt.Println(inflector.Camelize("multiple_http_calls"))
	// Output:
	// MyAccount
	// UserProfile
	// SSLError
	// HTTPConnectionTimeout
	// RESTfulController
	// MultipleHTTPCalls
}

func ExampleTitleize() {
	fmt.Println(inflector.Titleize("my_account"))
	fmt.Println(inflector.Titleize("user-profile"))
	fmt.Println(inflector.Titleize("ssl_error"))
	fmt.Println(inflector.Titleize("http_connection_timeout"))
	fmt.Println(inflector.Titleize("restful_controller"))
	// Output:
	// My Account
	// User Profile
	// SSL Error
	// HTTP Connection Timeout
	// RESTful Controller
}

func ExampleUnderscorize() {
	fmt.Println(inflector.Underscorize("MyAccount"))
	fmt.Println(inflector.Underscorize("user-profile"))
	fmt.Println(inflector.Underscorize("SSLError"))
	fmt.Println(inflector.Underscorize("HTTPConnectionTimeout"))
	fmt.Println(inflector.Underscorize("RESTfulController"))
	fmt.Println(inflector.Underscorize("MultipleHTTPCalls"))
	// Output:
	// my_account
	// user_profile
	// ssl_error
	// http_connection_timeout
	// restful_controller
	// multiple_http_calls
}

func ExampleDasherize() {
	fmt.Println(inflector.Dasherize("MyAccount"))
	fmt.Println(inflector.Dasherize("user_profile"))
	// Output:
	// my-account
	// user-profile
}

func ExampleTableize() {
	fmt.Println(inflector.Tableize("RawScaledScorer"))
	fmt.Println(inflector.Tableize("ham_and_egg"))
	fmt.Println(inflector.Tableize("fancyCategory"))
	// Output:
	// raw_scaled_scorers
	// ham_and_eggs
	// fancy_categories
}

func ExampleForeignKey() {
	fmt.Println(inflector.ForeignKey("Message"))
	fmt.Println(inflector.ForeignKey("AdminPost"))
	// Output:
	// message_id
	// admin_post_id
}

func ExampleOrdinal() {
	fmt.Println(inflector.Ordinal(1))
	fmt.Println(inflector.Ordinal(2))
	fmt.Println(inflector.Ordinal(14))
	fmt.Println(inflector.Ordinal(1002))
	fmt.Println(inflector.Ordinal(1003))
	fmt.Println(inflector.Ordinal(-11))
	fmt.Println(inflector.Ordinal(-1021))
	// Output:
	// st
	// nd
	// th
	// nd
	// rd
	// th
	// st
}

func ExampleOrdinalize() {
	fmt.Println(inflector.Ordinalize(1))
	fmt.Println(inflector.Ordinalize(2))
	fmt.Println(inflector.Ordinalize(14))
	fmt.Println(inflector.Ordinalize(1002))
	fmt.Println(inflector.Ordinalize(1003))
	fmt.Println(inflector.Ordinalize(-11))
	fmt.Println(inflector.Ordinalize(-1021))
	// Output:
	// 1st
	// 2nd
	// 14th
	// 1002nd
	// 1003rd
	// -11th
	// -1021st
}
