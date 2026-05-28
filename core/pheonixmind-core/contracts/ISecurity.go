/* Version: 0.1.0 */
package contracts
type ISecurity interface {
    Authorize(action string) bool
}
