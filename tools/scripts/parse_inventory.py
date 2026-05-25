import csv
import json

domains = {
    'Runtime Core': ['PheonixOS', 'ledger-core', 'control-plane', 'modules'],
    'Security': ['aegis-auth', 'SecureForg', 'TrustLab', 'smart-api-limiter'],
    'Infrastructure': ['astraeus-core', 'infrastructure'],
    'Intelligence / Agents': ['AI-PFI', 'ChoreoAI', 'Noesis', 'agentskill', 'cognitron-game', 'gametrend-intelligence-engine', 'forge-agent', 'brain', 'repo-analyzer'],
    'Science / Physics / Simulation': ['ParticleStimulator', 'physics', 'simulation'],
    'Health / Bio': ['AI4MH', 'TerraHerb', 'healingstone', 'LifeTrack', 'autoeit-suite', 'AutoEIT-STS', 'UDIE'],
    'Automation': ['AutoMation-Engine', 'AutoTRandHD'],
    'Products / UX': ['ArtExtract', 'LAMP', 'audio_transcription', 'my-portfolio', 'sira', 'idea', 'truenotes'],
    'Archive / Legacy': ['legacy']
}

def get_domain(repo_name):
    for domain, repos in domains.items():
        if repo_name in repos:
            return domain
    return 'Unknown'

with open('raw.csv', 'r') as infile, open('ecosystem/inventory/portfolio_inventory.csv', 'w', newline='') as outfile:
    writer = csv.writer(outfile)
    writer.writerow(['repo_name', 'domain', 'status', 'language', 'updated_at', 'is_archived'])
    for line in infile:
        parts = line.strip().split(',')
        if len(parts) == 4:
            name, lang, updated, archived = parts
            domain = get_domain(name)
            status = 'PENDING'
            writer.writerow([name, domain, status, lang, updated, archived])
