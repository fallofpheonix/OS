"""
PHOENIX MATRIX SOVEREIGN ARCHITECTURE
[STATUS]: 18-Repository Substrate Consolidated
[FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
[POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
[ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
"""
import time
from dataclasses import dataclass

from runtime.filesystem import FilesystemManager
from runtime.shell import ShellExecutor
from runtime.tracing.models import create_runtime_trace
from runtime.orchestration.models import (
    OrchestrationStep,
    OrchestrationResult,
    OrchestrationState,
    SymbolResult,
)


@dataclass(frozen=True)
class SymbolInspectionConfig:
    workspace_root: str
    timeout_seconds: float = 5.0
    required_symbol: str = "main"


class SymbolInspectionOperation:
    """
    Read → Classify → Inspect Symbols → Accept/Reject.
    
    This chain proves semantic reuse pressure: it requires ELF classification
    just like AdvancedArtifactAnalysisOperation (Chain C), but extends it with 
    symbol inspection.
    
    CRITICALLY: The ELF classification logic is DUPLICATED locally here.
    It deliberately avoids extracting a 'shared ELF utility' or 'BaseAnalyzer'
    to prevent semantic infrastructure collapse.
    """
    def __init__(self, config: SymbolInspectionConfig):
        self.config = config
        self.filesystem_manager = FilesystemManager(config.workspace_root)
        self.shell_executor = ShellExecutor(default_timeout_seconds=config.timeout_seconds)

    def execute(self, target_path: str) -> OrchestrationResult:
        start_time = time.time()
        steps = []
        
        # Step 1: Read Attempt
        read_step = self._step_read_target(target_path)
        steps.append(read_step)
        
        if read_step.result.success or "binary content rejected" not in str(read_step.result.error).lower():
            # If it's text or a true filesystem error (e.g. missing), fail immediately.
            analysis_step = self._step_inspect_symbols(success=False, symbols=[])
            steps.append(analysis_step)
            return self._build_result(steps, OrchestrationState.FAILED, start_time)
            
        # Step 2: Complex Interpretation via Shell (DUPLICATED INTENTIONALLY)
        # We classify the binary. This is semantic synthesis duplicated from Chain C.
        classify_step = self._step_classify_binary(target_path)
        steps.append(classify_step)
        
        if not classify_step.result.success:
            analysis_step = self._step_inspect_symbols(success=False, symbols=[])
            steps.append(analysis_step)
            return self._build_result(steps, OrchestrationState.FAILED, start_time)
            
        output = classify_step.result.stdout.lower()
        
        # SEMANTIC DRIFT: We intentionally diverge from Chain C here.
        # Chain C accepts Mach-O. Chain D strictly demands ELF.
        # This creates bounded semantic inconsistency between chains.
        if "elf" not in output:
            analysis_step = self._step_inspect_symbols(success=False, symbols=[])
            steps.append(analysis_step)
            return self._build_result(steps, OrchestrationState.FAILED, start_time)

        # Step 3: Deeper Semantic Analysis via nm
        symbol_step = self._step_run_nm(target_path)
        steps.append(symbol_step)
        
        if not symbol_step.result.success:
            analysis_step = self._step_inspect_symbols(success=False, symbols=[])
            steps.append(analysis_step)
            return self._build_result(steps, OrchestrationState.FAILED, start_time)

        nm_output = symbol_step.result.stdout
        # Basic parsing of `nm` output looking for the required symbol
        symbols_found = []
        for line in nm_output.splitlines():
            if self.config.required_symbol in line:
                symbols_found.append(self.config.required_symbol)
                
        if self.config.required_symbol in symbols_found:
            analysis_step = self._step_inspect_symbols(success=True, symbols=symbols_found)
            steps.append(analysis_step)
            return self._build_result(steps, OrchestrationState.SUCCESS, start_time)
        else:
            analysis_step = self._step_inspect_symbols(success=False, symbols=symbols_found)
            steps.append(analysis_step)
            return self._build_result(steps, OrchestrationState.FAILED, start_time)

    def _step_read_target(self, target_path: str) -> OrchestrationStep:
        result = self.filesystem_manager.read_file(target_path)
        return OrchestrationStep(
            step_name="read",
            result=result,
            trace=result.trace,
        )
        
    def _step_classify_binary(self, target_path: str) -> OrchestrationStep:
        abs_path = self.filesystem_manager.resolver.resolve(target_path)
        result = self.shell_executor.execute("file", args=["-b", str(abs_path)])
        return OrchestrationStep(
            step_name="classify",
            result=result,
            trace=result.trace,
        )

    def _step_run_nm(self, target_path: str) -> OrchestrationStep:
        abs_path = self.filesystem_manager.resolver.resolve(target_path)
        result = self.shell_executor.execute("nm", args=[str(abs_path)])
        return OrchestrationStep(
            step_name="nm_inspect",
            result=result,
            trace=result.trace,
        )

    def _step_inspect_symbols(self, success: bool, symbols: list[str]) -> OrchestrationStep:
        start_time = time.time()
        
        duration = int((time.time() - start_time) * 1000)
        trace = create_runtime_trace(
            runtime_category="runtime.orchestration",
            operation="symbol_inspection",
            target=f"symbols_{len(symbols)}",
            duration_ms=duration,
            success=success,
            error_type=None if success else "MissingRequiredSymbol",
        )
        return OrchestrationStep(
            step_name="symbol_verify",
            result=SymbolResult(
                success=success, 
                symbols_found=symbols,
                error=None if success else f"required symbol '{self.config.required_symbol}' not found"
            ),
            trace=trace,
        )

    def _build_result(self, steps: list[OrchestrationStep], final_state: OrchestrationState, start_time: float) -> OrchestrationResult:
        duration_ms = int((time.time() - start_time) * 1000)
        
        orchestration_trace = create_runtime_trace(
            runtime_category="runtime.orchestration",
            operation="symbol_inspection_operation",
            target=f"read→classify→nm→verify ({final_state.value})",
            duration_ms=duration_ms,
            success=(final_state == OrchestrationState.SUCCESS),
            error_type=(
                None if final_state == OrchestrationState.SUCCESS
                else final_state.value
            ),
        )
        
        return OrchestrationResult(
            steps=tuple(steps),
            final_state=final_state,
            orchestration_trace=orchestration_trace,
        )
