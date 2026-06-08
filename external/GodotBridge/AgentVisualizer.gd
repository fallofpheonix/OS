extends Node3D

# Godot WebSocket Bridge for Phoenix Game Layer
# Connects to Go VM and visualizes world state.

var socket = WebSocketPeer.new()
var url = "ws://localhost:8080/ws"
var precision = 1000000.0

@onready var agent_node = $AgentMesh # Assumes a MeshInstance3D child named AgentMesh
@onready var status_label = $StatusLabel # Assumes a Label3D child named StatusLabel

func _ready():
	socket.connect_to_url(url)
	print("Connecting to Phoenix Server at ", url)

func _process(_delta):
	socket.poll()
	var state = socket.get_ready_state()
	
	if state == WebSocketPeer.STATE_OPEN:
		while socket.get_available_packet_count() > 0:
			var packet = socket.get_packet()
			_handle_world_state(packet.get_string_from_utf8())
	elif state == WebSocketPeer.STATE_CLOSED:
		var code = socket.get_close_code()
		var reason = socket.get_close_reason()
		print("WebSocket closed with code: %d, reason %s. Reconnecting..." % [code, reason])
		socket.connect_to_url(url)

func _handle_world_state(json_str: String):
	var json = JSON.new()
	var error = json.parse(json_str)
	if error != OK:
		print("JSON Parse Error: ", json.get_error_message())
		return
	
	var data = json.get_data()
	var tick = data.get("tick", 0)
	var entities = data.get("entities", {})
	
	if entities.has("agent_01"):
		var agent = entities["agent_01"]
		var raw_pos = agent.get("pos", {}).get("value", 0)
		var float_pos = float(raw_pos) / precision
		var status = agent.get("status", "UNKNOWN")
		
		# Update Visuals
		if agent_node:
			agent_node.position.x = float_pos
		
		if status_label:
			status_label.text = "Tick: %d\nPos: %.2f\nStatus: %s" % [tick, float_pos, status]
			
		print("Tick %d | Agent Pos: %.2f | Status: %s" % [tick, float_pos, status])
