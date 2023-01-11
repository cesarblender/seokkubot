package lib

import (
	"fmt"
	"log"
	"mybot/settings"
	"strconv"

	fb "github.com/huandu/facebook/v2"
)

func CreatePost(text string, likes int, omitted_comments int, image string, createdTime string) {
	_, err := fb.Post("/me/photos", fb.Params{
		"message": text + `

⏱ Creado el: ` + createdTime + `
✔ Votaciones positivas: ` + strconv.Itoa(int(likes)) + `
⛔ Comentarios omitidos: ` + strconv.Itoa(int(omitted_comments)) + `

¿Cómo puedo publicar en el bot?

Comenta la última publicación del bot, en el siguiente formato

!publicame
<Aquí colocas el texto de arriba de la publicación, (sin incluir los "<>" xd)>

Además de comentar eso, debes agregar una imágen para que tu solicitud sea válida.

¿Cómo puedo votar por un comentario?
El bot, en la última publicación selecciona el comentario más votado para publicarlo, y para votar por un comentario, hay que darle like. Sólo los likes son contados como votos positivos, no se toma en cuenta ninguna otra reacción, debido a políticas de Facebook.

¿Por qué no aparecen las publicaciónes que solicite?

La razon de esto pueden ser varias.

- La primera y más común sería que el bot omitió tu comentario, ya que había otro comentario con mas likes, el bot automáticamente selecciona el comentario con mas likes.

- La segunda razón, puede ser que ya haya una publicación mas reciente, al haber otra publicación mas reciente, automáticamente se ignorarán los comentarios de las publicaciones anteriores. Se publica una nueva publicación 30 minutos después de la última publicación.

- La tercera, puede ser que hayas escrito mal "!publicame" el bot está configurado para que detecte si un comentario comienza con publicame | publícame | publìcame | publocame | publicamee | piblicame | poblicame, para evitar este tipo de errores que son muy comunes. También cabe recalcar, que tu comentario debe comenzar con un "!" nunca con otro caracter, es decir, si escribiste ",!publicame" automáticamente va a ser ignorado tu comentario, ya que no coincide con el requisito para que sea válido.

- La cuarta razón puede ser que el bot esté en mantenimiento y no haya detectado tu comentario.

Autor del bot: seokku papu え
Código fuente del bot: https://github.com/seokkuuu/seokkubot
GitHub del autor del bot: https://github.com/seokkuuu
Proyecto bajo la licencia copyleft: (CC) Creative Commons Attribution

Ningún post está relacionado con el autor del bot, todas las publicaciones son realizadas a partir de comentarios que realizan usuarios, el bot automáticamente selecciona el más votado y lo publica.

Tags: #seokkubot #randomcatbot #bot #shitposting #momos #memes #sdlg #grasa #curifeos #curipapus #papus #go #golang #github`,
		"url":          image,
		"access_token": settings.GetEnv("ACCESS_TOKEN"),
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Post created")
}
