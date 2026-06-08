/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
const express = require("express");
const router = express.Router();
const {
  createNote,
  getAllNotes,
  searchNotes,
  getNote,
  updateNote,
  deleteNote,
} = require("../controllers/noteController");
const authMiddleware = require("../middlewares/authMiddleware");

router.use(authMiddleware);

router.post("/", createNote);
router.get("/", getAllNotes);
router.get("/search", searchNotes);
router.get("/:id", getNote);
router.put("/:id", updateNote);
router.delete("/:id", deleteNote);

module.exports = router;
