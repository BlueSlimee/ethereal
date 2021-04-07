import pylast
import os
import requests
from urllib.parse import urlencode


def get_top_albums(user, limit=9, p='overall'):
    r = request('user.getTopAlbums', data={
        'user': user,
        'period': p,
        'limit': limit,
    })
    
    if r == None:
        return None
    
    albums = r['topalbums']['album']
    return list(map(lambda x: {
        'artist_name': x['artist']['name'],
        'scrobbles': x['playcount'],
        'name': x['name'],
        'image': x['image'][2]['#text'] if len(x['image']) > 2 else 'https://lastgram.vercel.app/last/missingalbum.jpg'
    }, albums))

def request(method, data={}):
    headers = {
        'User-Agent': 'Lastgram Ethereal'
    }

    data['method'] = method
    data['api_key'] = os.getenv('FM_API_KEY')
    data['format'] = 'json'

    r = requests.get('http://ws.audioscrobbler.com/2.0/', params=data, headers=headers)
    return r.json()
