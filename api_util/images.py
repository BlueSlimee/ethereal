from io import BytesIO
from PIL import Image, ImageFont
import requests

def to_jpeg(img):
    io = BytesIO()
    img.save(io, 'JPEG', quality=75)
    io.seek(0)
    return io

def get_font(family, sec, size):
    return ImageFont.truetype(join(
      'public',
        'fonts',
        family,
        sec + '.ttf'
    ), size)

def get_image_from_url(url):
    r = requests.get(url)
    return Image.open(BytesIO(r.content)) if r.status_code == 200 else None
