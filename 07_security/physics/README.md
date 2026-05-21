# Incident Physics (L6)

Statistical mechanics engine for Pheonix security state estimation.

## Purpose
Model system-wide security as a thermodynamic process to identify cascading failure points.

## Formulas
### Security Disorder Index (SDI)
$$\theta_{\text{SDI}} = -\sum p_s \ln p_s$$

### Hamiltonian (Ising Model)
$$H(\sigma) = -J \sum \sigma_i \sigma_j - h \sum \sigma_i$$

## Performance Budget
- **SDI Calculation:** < 100 us.

## Validation Gates
- [ ] Correctness of entropy calculation on microstates.
- [ ] Phase transition prediction accuracy.
