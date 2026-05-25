/* Version: 0.1.0 */
package contracts
type ISandbox interface {
    Run(cmd string) error
}
