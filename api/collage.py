from flask import Flask, Response, send_file
from PIL import Image, ImageDraw, ImageFont
from os.path import join
from io import BytesIO
app = Flask(__name__)

def gen_collage(w, h):
    img = Image.new('RGB', (w * 100, h * 100), color = 'black')
    fnt = get_font('montserrat', 'bold', 30)
    draw = ImageDraw.draw(img)
    draw.text((10, 10), "paguei um boquete pra jesus", font=fnt, fill=(0, 0, 0))

    return to_jpeg(img)

@app.route('/', defaults={'path': ''})
@app.route('/<path:path>')
def root(path):
    return send_file(gen_collage(request.args.get('w'), request.args.get('h')), mimetype='image/jpeg')

def to_jpeg(img):
    io = BytesIO()
    img.save(io, 'JPEG', quality=75)
    io.seek(0)
    return io

def get_font(family, sec, size):
    return ImageFont.truetype(join('public', 'fonts', family, sec +'.ttf'), size)
