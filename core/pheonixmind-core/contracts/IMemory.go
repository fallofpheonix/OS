/* Version: 0.1.0 */
package contracts
type IMemory interface {
    Recall(key string) ([]byte, error)
}
