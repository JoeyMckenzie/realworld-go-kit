package comments

import (
    "context"
    "database/sql"
    "github.com/google/uuid"
    "github.com/joeymckenzie/realworld-go-kit/internal/shared"
    "net/http"
)

func (cs *commentsService) DeleteComment(ctx context.Context, userId, commentId uuid.UUID) error {
    // First, verify the comment exist
    existingComment, err := cs.commentsRepository.GetArticleComment(ctx, commentId, userId)

    if shared.IsValidSqlErr(err) {
        cs.logger.ErrorCtx(ctx, "error while attempting to check for comment before deleting", "user_id", userId, "comment_id", commentId)
        return shared.MakeApiError(err)
    } else if err == sql.ErrNoRows {
        cs.logger.ErrorCtx(ctx, "comment not found", "user_id", userId, "comment_id", commentId)
        return shared.MakeApiErrorWithStatus(http.StatusNotFound, shared.ErrCommentNotFound)
    }

    // Next, verify the user owns the comments
    if existingComment.AuthorId != userId {
        cs.logger.ErrorCtx(ctx, "user attempting to remove comment they did not author", "user_id", userId, "author_id", existingComment.AuthorId, "comment_id", commentId)
        // Similar to other forbidden accesses, return a not found to discourage probing attacks
        return shared.MakeApiErrorWithStatus(http.StatusNotFound, shared.ErrCommentNotFound)
    }

    // Finally, delete the comment
    return cs.commentsRepository.DeleteComment(ctx, commentId)
}
