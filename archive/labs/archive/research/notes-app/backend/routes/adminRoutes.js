/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
const express = require("express");
const router = express.Router();
const adminController = require("../controllers/adminController");
const authMiddleware = require("../middlewares/authMiddleware");
const isAdmin = require("../middlewares/adminMiddleware");

// Apply authentication and admin check middleware to all routes
router.use(authMiddleware);
router.use(isAdmin);

// User management
router.get("/users", adminController.getAllUsers);
router.get("/users/:id", adminController.getUserById);
router.put("/users/:id/role", adminController.updateUserRole);
router.delete("/users/:id", adminController.deleteUser);

// Note management
router.get("/notes", adminController.getAllNotes);

// Statistics
router.get("/stats", adminController.getUserStats);

module.exports = router;