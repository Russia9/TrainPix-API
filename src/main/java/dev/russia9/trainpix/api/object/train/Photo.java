package dev.russia9.trainpix.api.object.train;

public class Photo {
    /**
     * Trainpix Photo ID
     * <p>
     * Example:
     * https://trainpix.org/photo/266009/ - 266009
     */
    private int id;

    /**
     * Image URI
     * <p>
     * Example:
     * https://trainpix.org/photo/02/66/00/266009.jpg
     */
    private String image;

    /**
     * Thumbnail URI
     * <p>
     * Example:
     * https://trainpix.org/photo/02/66/00/266009_s.jpg
     */
    private String thumbnail;

    /**
     * Image page URI
     * <p>
     * Example:
     * https://trainpix.org/photo/266009
     */
    private String page;

    /**
     * Image size
     * <p>
     * Example:
     * https://trainpix.org/photo/266009 - 680 KB
     */
    private String size;

    /**
     * Image date
     * <p>
     * Example:
     * https://trainpix.org/photo/181451/ - December 4, 2016
     */
    private String date;

    /**
     * Image publication date
     * <p>
     * Example:
     * https://trainpix.org/photo/181451/ - 04.12.2016 20:30 EET
     */
    private String publishedDate;

    /**
     * Image author name
     * <p>
     * Example:
     * https://trainpix.org/photo/181451/ - Максим П.
     */
    private String author;

    /**
     * Image author link
     * <p>
     * Example:
     * https://trainpix.org/photo/181451/ - https://trainpix.org/author/4463/
     */
    private String authorLink;

    public Photo(int id) {
        this.id = id;
    }

    public void setImage(String image) {
        this.image = image;
    }

    public void setThumbnail(String thumbnail) {
        this.thumbnail = thumbnail;
    }

    public void setPage(String page) {
        this.page = page;
    }

    public void setSize(String size) {
        this.size = size;
    }

    public void setDate(String date) {
        this.date = date;
    }

    public void setPublishedDate(String publishedDate) {
        this.publishedDate = publishedDate;
    }

    public void setAuthor(String author) {
        this.author = author;
    }

    public void setAuthorLink(String authorLink) {
        this.authorLink = authorLink;
    }
}
