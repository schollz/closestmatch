package test

import (
	"strings"
)

var books = `Pride and Prejudice by Jane Austen
Alice's Adventures in Wonderland by Lewis Carroll
The Importance of Being Earnest: A Trivial Comedy for Serious People by Oscar Wilde
A Tale of Two Cities by Charles Dickens
A Doll's House : a play by Henrik Ibsen
Frankenstein; Or, The Modern Prometheus by Mary Wollstonecraft Shelley
The Yellow Wallpaper by Charlotte Perkins Gilman
The Adventures of Tom Sawyer by Mark Twain
Metamorphosis by Franz Kafka
Adventures of Huckleberry Finn by Mark Twain
Light Science for Leisure Hours by Richard A. Proctor
Grimms' Fairy Tales by Jacob Grimm and Wilhelm Grimm
Jane Eyre: An Autobiography by Charlotte Brontë
Dracula by Bram Stoker
Moby Dick; Or, The Whale by Herman Melville
The Adventures of Sherlock Holmes by Arthur Conan Doyle
Il Principe. English by Niccolò Machiavelli
Emma by Jane Austen
Great Expectations by Charles Dickens
The Picture of Dorian Gray by Oscar Wilde
Beyond the Hills of Dream by W. Wilfred Campbell
The Hospital Murders by Means Davis and Augusta Tucker Townsend
Dirty Dustbins and Sloppy Streets by H. Percy Boulnois
Leviathan by Thomas Hobbes
The Count of Monte Cristo, Illustrated by Alexandre Dumas
Heart of Darkness by Joseph Conrad
Ulysses by James Joyce
War and Peace by graf Leo Tolstoy
Narrative of the Life of Frederick Douglass, an American Slave by Frederick Douglass
The Radio Boys Seek the Lost Atlantis by Gerald Breckenridge
The Bab Ballads by W. S. Gilbert
Wuthering Heights by Emily Brontë
The Awakening, and Selected Short Stories by Kate Chopin
The Romance of Lust: A Classic Victorian erotic novel by Anonymous
Beowulf
Les Misérables by Victor Hugo
Siddhartha by Hermann Hesse
The Kama Sutra of Vatsyayana by Vatsyayana
Treasure Island by Robert Louis Stevenson
Dubliners by James Joyce
Reminiscences of Western Travels by Shao Xiang Lin
The Souls of Black Folk by W. E. B. Du Bois
Leaves of Grass by Walt Whitman
A Christmas Carol in Prose; Being a Ghost Story of Christmas by Charles Dickens
Tractatus Logico-Philosophicus by Ludwig Wittgenstein
A Modest Proposal by Jonathan Swift
Essays of Michel de Montaigne — Complete by Michel de Montaigne
Prestuplenie i nakazanie. English by Fyodor Dostoyevsky
Practical Grammar and Composition by Thomas Wood
A Study in Scarlet by Arthur Conan Doyle
Sense and Sensibility by Jane Austen
Don Quixote by Miguel de Cervantes Saavedra
Peter Pan by J. M. Barrie
The Republic by Plato
The Life and Adventures of Robinson Crusoe by Daniel Defoe
The Strange Case of Dr. Jekyll and Mr. Hyde by Robert Louis Stevenson
Gulliver's Travels into Several Remote Nations of the World by Jonathan Swift
My Secret Life, Volumes I. to III. by Anonymous
Beyond Good and Evil by Friedrich Wilhelm Nietzsche
The Brothers Karamazov by Fyodor Dostoyevsky
The Time Machine by H. G. Wells
Also sprach Zarathustra. English by Friedrich Wilhelm Nietzsche
The Federalist Papers by Alexander Hamilton and John Jay and James Madison
Songs of Innocence, and Songs of Experience by William Blake
The Iliad by Homer
Hastings & Environs; A Sketch-Book by H. G. Hampton
The Hound of the Baskervilles by Arthur Conan Doyle
The Children of Odin: The Book of Northern Myths by Padraic Colum
Autobiography of Benjamin Franklin by Benjamin Franklin
The Divine Comedy by Dante, Illustrated by Dante Alighieri
Hedda Gabler by Henrik Ibsen
Hard Times by Charles Dickens
The Jungle Book by Rudyard Kipling
The Real Captain Kidd by Cornelius Neale Dalton
On Liberty by John Stuart Mill
The Complete Works of William Shakespeare by William Shakespeare
The Tragical History of Doctor Faustus by Christopher Marlowe
Anne of Green Gables by L. M. Montgomery
The Jungle by Upton Sinclair
The Tragedy of Romeo and Juliet by William Shakespeare
De l'amour by Charles Baudelaire and Félix-François Gautier
Ethan Frome by Edith Wharton
Oliver Twist by Charles Dickens
The Turn of the Screw by Henry James
The Wonderful Wizard of Oz by L. Frank Baum
The Legend of Sleepy Hollow by Washington Irving
The Ship of Coral by H. De Vere Stacpoole
Democracy and Education: An Introduction to the Philosophy of Education by John Dewey
Candide by Voltaire
Pygmalion by Bernard Shaw
Walden, and On The Duty Of Civil Disobedience by Henry David Thoreau
Three Men in a Boat by Jerome K. Jerome
A Portrait of the Artist as a Young Man by James Joyce
Manifest der Kommunistischen Partei. English by Friedrich Engels and Karl Marx
Through the Looking-Glass by Lewis Carroll
Le Morte d'Arthur: Volume 1 by Sir Thomas Malory
The Mysterious Affair at Styles by Agatha Christie
Korean—English Dictionary by Leon Kuperman
The War of the Worlds by H. G. Wells
A Concise Dictionary of Middle English from A.D. 1150 to 1580 by A. L. Mayhew and Walter W. Skeat
Armageddon in Retrospect by Kurt Vonnegut
Red Riding Hood by Sarah Blakley-Cartwright
The Kingdom of This World by Alejo Carpentier
Hitty, Her First Hundred Years by Rachel Field`

var WordsToTest []string
var SearchWords = []string{"cervantes don quixote", "mysterious afur at styles by christie", "hard times by charles dickens", "complete william shakespeare", "War by HG Wells"}

func init() {
	WordsToTest = strings.Split(strings.ToLower(books), "\n")
	for i := range SearchWords {
		SearchWords[i] = strings.ToLower(SearchWords[i])
	}
}
