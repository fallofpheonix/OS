# Module: pfi-ingestion

## Source
Extracted from [[05_PROJECTS/COMPLETED/AI-PFI/Project]]

## Purpose
Web scraping → data normalization → ontology tagging pipeline for ingesting structured data from external sources (government portals, APIs, web pages).

## Interface
```python
from pfi_ingestion import Scraper, Normalizer, OntologyTagger, Pipeline

# Individual components
scraper = Scraper(sources=["grants.gov", "nsf.gov"])
raw_data = scraper.fetch(query="machine learning")

normalizer = Normalizer(schema=FOASchema)
clean_data = normalizer.transform(raw_data)

tagger = OntologyTagger(ontology="funding_categories")
tagged_data = tagger.tag(clean_data)

# Or as a pipeline
pipeline = Pipeline(scraper, normalizer, tagger)
results = pipeline.run(query="AI safety research")
```

## Depends On
- [[05_PROJECTS/REUSABLE_MODULES/fpx-pipeline]] (pipeline control-plane)
- requests/httpx (HTTP client)
- BeautifulSoup/selectolax (HTML parsing)

## Used By
- Personal Knowledge Graph (web content ingestion)
- Banking App (regulatory data feeds, market data)

## Extraction Status
NOT_STARTED

## Location
`~/engineering/infrastructure/shared-libraries/pfi-ingestion/`

## Key Files
| File | Role |
|------|------|
| `scraper.py` | Configurable web scraper with rate limiting |
| `normalizer.py` | Schema-driven data normalization |
| `tagger.py` | Lightweight ontology tagging engine |
| `pipeline.py` | Compose scraper → normalizer → tagger into a single pipeline |
| `schemas/` | Pydantic models for different data source formats |

## Quality Gates
- [ ] Tests passing
- [ ] Scraper respects robots.txt
- [ ] Normalizer is schema-driven (not hardcoded to Grants.gov)
- [ ] README with example pipeline for a new data source
- [ ] Version pinned

#module #extracted-from/AI-PFI #priority/P3
