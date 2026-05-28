# Model Routing Rules

architecture:
  model: GeminiCLI
  write_access: false

code_generation:
  model: local_coder
  write_access: sandbox_only

security:
  model: validator_llm
  write_access: none

merge:
  model: human
  write_access: production
