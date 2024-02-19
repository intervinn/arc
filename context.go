package arc

import "encoding/json"

// parse json body, argument must be a pointer to structure
func (c *Ctx) BodyParse(v any) error {
	return json.NewDecoder(c.Request.Body).Decode(v)
}

// send json response
func (c *Ctx) JSON(status int, v any) error {
	c.Response.WriteHeader(status)
	c.Response.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(c.Response).Encode(v)
}

// write bytes to response
func (c *Ctx) Write(status int, bytes []byte) (int, error) {
	c.Response.WriteHeader(status)
	return c.Response.Write(bytes)
}

// write string to respones
func (c *Ctx) WriteString(status int, str string) (int, error) {
	return c.Write(status, []byte(str))
}

// get value from request-exclusive map, string is empty is value doesnt exist
func (c *Ctx) Get(key string) any {
	return c.local[key]
}

// set value for request-exclusive map
func (c *Ctx) Set(key string, val any) {
	c.local[key] = val
}
