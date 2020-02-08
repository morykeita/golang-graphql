/**
 * @author Mory Keita on 2/7/20
 */
package models

type User struct {
	ID       string    `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Meetups  []*Meetup `json:"meetups"`
}
