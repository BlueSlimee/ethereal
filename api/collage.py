from flask import Flask, Response, send_file, request
from PIL import Image, ImageDraw, ImageFont
from os.path import join
from io import BytesIO
from api_util.images import to_jpeg, get_font, get_image_from_url
from api_util.lfm import get_top_albums

app = Flask(__name__)

def gen_collage(u, w, h):
    total = w * h
    mw = w * 174
    mh = h * 174
    
    if total > 100:
        return None
    
    img = Image.new('RGB', (mw, mh))
    #fnt = get_font('montserrat', 'bold', 30)
    #draw = ImageDraw.Draw(img)
    #draw.text((0, 0), "paguei um boquete pra jesus", align='center', font=fnt, fill='white')
    currentx = 0
    currenty = 0
    lista = get_top_albums(u, limit=total)
    i = 0
    
    while i < total:
        item = lista[i] 
        img.paste(
            get_image_from_url(item['image']),
            (currentx, currenty)
        )
        currentx = currentx + 174
        if currentx >= mw:
            currentx = 0
            currenty = currenty + 174

        i = i + 1
    return to_jpeg(img)

@app.route('/', defaults={'path': ''})
@app.route('/<path:path>', methods=['GET'])
def root(path):
    if request.args.get('r') == 'collage':
        return send_file(gen_collage(
            request.args.get('u'),
            int(request.args.get('w')),
            int(request.args.get('h'))
        ), mimetype='image/jpeg')

