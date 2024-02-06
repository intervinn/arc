package arc

import "encoding/json"

func (c *Ctx) BodyParse(v any) error {
	return json.NewDecoder(c.Request.Body).Decode(v)
}

func (c *Ctx) JSON(status int, v any) error {
	c.Response.WriteHeader(status)
	c.Response.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(c.Response).Encode(v)
}

func (c *Ctx) Write(status int, bytes []byte) (int, error) {
	c.Response.WriteHeader(status)
	return c.Response.Write(bytes)
}

func (c *Ctx) WriteString(status int, str string) (int, error) {
	c.Response.WriteHeader(status)
	return c.Response.Write([]byte(str))
}
