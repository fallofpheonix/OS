/* Version: 0.1.0 */
package contracts
type IValidator interface {
    Validate(data []byte) (bool, error)
}
