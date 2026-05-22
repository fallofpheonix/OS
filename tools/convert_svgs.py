#!/usr/bin/env python3
"""Convert SVGs in assets/pheonix to PNG and ICO outputs.

Tries cairosvg, falls back to creating simple PNG placeholders if unavailable.
"""
from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]
ASSETS = ROOT / 'assets' / 'pheonix'
OUT = ASSETS / 'raster'
OUT.mkdir(parents=True, exist_ok=True)

svgs = list(ASSETS.glob('*.svg'))
if not svgs:
    print('No SVGs found in', ASSETS)
    raise SystemExit(1)

try:
    import cairosvg
    have_cairosvg = True
except Exception:
    have_cairosvg = False

for s in svgs:
    name = s.stem
    png_out = OUT / f'{name}.png'
    ico_out = OUT / f'{name}.ico'
    if have_cairosvg:
        print('Converting', s, '->', png_out)
        cairosvg.svg2png(url=str(s), write_to=str(png_out), output_width=512, output_height=512)
        # create ico with same png data (simple approach)
        try:
            from PIL import Image
            im = Image.open(png_out)
            im.save(ico_out)
        except Exception:
            # fallback: copy png to ico name
            if not ico_out.exists():
                open(ico_out, 'wb').write(open(png_out,'rb').read())
        print('WROTE', png_out, ico_out)
    else:
        print('cairosvg not available; creating placeholder PNG for', name)
        # write a tiny 1x1 PNG fallback (transparent)
        import base64
        tiny_png_b64 = 'iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR4nGNgYAAAAAMAASsJTYQAAAAASUVORK5CYII='
        png_bytes = base64.b64decode(tiny_png_b64)
        with open(png_out, 'wb') as f:
            f.write(png_bytes)
        with open(ico_out, 'wb') as f:
            f.write(png_bytes)
        print('WROTE tiny placeholder', png_out, ico_out)

print('Done')
